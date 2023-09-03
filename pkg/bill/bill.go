package bill

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/utils"
)

//go:generate mockgen -destination=./mocks/mock_bill.go -package=mock_bill github.com/sanchitdeora/budget-tracker/pkg/bill Service
type Service interface {
	GetBills(ctx context.Context) (*[]models.Bill, error)
	GetBillById(ctx context.Context, id string) (*models.Bill, error)
	CreateBillByUser(ctx context.Context, bill *models.Bill) (string, error)
	UpdateBillById(ctx context.Context, id string, bill *models.Bill) (string, error)
	UpdateBillIsPaid(ctx context.Context, id string) (string, error)
	UpdateBillIsUnpaid(ctx context.Context, id string) (string, error)
	DeleteBillById(ctx context.Context, id string) (string, error)

	// Maintain recurring bills
	BillMaintainer(ctx context.Context)
	CreateRecurringBill(ctx context.Context, bill *models.Bill) (string, error)
}

type Opts struct {
	TransactionService  transaction.Service
	DB					db.Database	
}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}


func (s *serviceImpl) GetBills(ctx context.Context) (*[]models.Bill, error) {
	return s.DB.GetAllBillRecords(ctx)
}

func (s *serviceImpl) GetBillById(ctx context.Context, id string) (*models.Bill, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return nil, exceptions.ErrValidationError
	}	
	return s.DB.GetBillRecordById(ctx, id)
}

func (s *serviceImpl) CreateBillByUser(ctx context.Context, bill *models.Bill) (string, error) {
	if !bill.IsValid() {
		return "", exceptions.ErrValidationError
	}

	bill.SetByUser()
	bill.SetCategory()
	bill.SetFrequency()
	return s.DB.InsertBillRecord(ctx, bill)
}

func (s *serviceImpl) CreateRecurringBill(ctx context.Context, prevBill *models.Bill) (string, error) {
	if !prevBill.IsValid() {
		return "", exceptions.ErrValidationError
	}

	var newBill models.Bill
	newBill.Title = prevBill.Title
	newBill.Category = prevBill.Category
	newBill.AmountDue = prevBill.AmountDue
	newBill.Frequency = prevBill.Frequency
	newBill.DueDate = utils.CalculateEndDateWithFrequency(prevBill.DueDate, prevBill.Frequency)
	newBill.IsPaid = false
	newBill.Note = prevBill.Note
	newBill.SequenceNumber = prevBill.SequenceNumber + 1

	if prevBill.SequenceStartId == "" {
		newBill.SequenceStartId = prevBill.BillId
	} else {
		newBill.SequenceStartId = prevBill.SequenceStartId
	}

	newBill.SetCategory()
	newBill.SetFrequency()
	return s.DB.InsertBillRecord(ctx, &newBill)
}

func (s *serviceImpl) UpdateBillById(ctx context.Context, id string, bill *models.Bill) (string, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return "", exceptions.ErrValidationError
	}

	if !bill.IsValid() {
		return "", exceptions.ErrValidationError
	}

	bill.SetCategory()
	bill.SetFrequency()
	return s.DB.UpdateBillRecordById(ctx, id, bill)
}

func (s *serviceImpl) UpdateBillIsPaid(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return "", exceptions.ErrValidationError
	}

	var bill *models.Bill
	bill, err := s.GetBillById(ctx, id)
	if err != nil {
		log.Println("error while fetching bill record for billId:", id, "error:", err, ctx)
		return "", err
	}

	if bill.IsPaid && bill.TransactionId != "" {
		return bill.BillId, nil
	}
	
	bill.IsPaid = true
	bill.DatePaid = time.Now()

	if bill.NextSequenceId == "" && bill.Frequency != models.ONCE_FREQUENCY {
		// create new bill entry for next frequency period
		newBillId, err := s.CreateRecurringBill(ctx, bill)
		if err != nil {
			log.Println("error while creating new bill id:", newBillId, "error:", err, ctx)
			return "", err
		}
		bill.NextSequenceId = newBillId
	}

	if bill.TransactionId == "" {
		var newTransaction models.Transaction
		newTransaction.FromBill(bill)
		transactionId, err := s.TransactionService.CreateTransaction(ctx, &newTransaction)
		if err != nil {
			log.Println("error while creating transaction record", newTransaction, "error:", err, ctx)
			return "", err
		}
		bill.TransactionId = transactionId
	} else {
		transaction, err := s.TransactionService.GetTransactionById(ctx, bill.TransactionId)
		if err != nil {
			log.Println("error while fetching transaction record for id", bill.TransactionId, "error:", err, ctx)
			return "", err
		}
		transaction.Date = bill.DatePaid
		_, err = s.TransactionService.UpdateTransactionById(ctx, transaction.TransactionId, transaction)
		if err != nil {
			log.Println("error while creating transaction record", transaction, "error:", err, ctx)
			return "", err
		}
	}

	return s.DB.UpdateBillRecordById(ctx, id, bill)
}

func (s *serviceImpl) UpdateBillIsUnpaid(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return "", exceptions.ErrValidationError
	}

	var bill *models.Bill
	bill, err := s.GetBillById(ctx, id)
	if err != nil {
		log.Println("error while fetching bill record for billId:", id, "error:", err, ctx)
		return "", err
	}

	if bill.TransactionId != "" {
		_, err := s.TransactionService.DeleteTransactionById(ctx, bill.TransactionId)
		if err != nil {
			log.Println("error while deleting transaction record for transactionId:", bill.TransactionId, "error:", err, ctx)
			return "", err
		}
		bill.TransactionId = ""
	}

	bill.DatePaid = time.Time{}
	bill.IsPaid = false

	return s.DB.UpdateBillRecordById(ctx, id, bill)
}

func (s *serviceImpl) DeleteBillById(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return "", exceptions.ErrValidationError
	}

	return s.DB.DeleteBillRecordById(ctx, id)
}


// Maintainer functions
func (s *serviceImpl) BillMaintainer(ctx context.Context) {
	log.Println("start bill maintainer", ctx)

	successfulCounter := 0
	bills, err := s.GetBills(ctx)
	if err != nil {
		log.Println("error while fetching bill records", "error:", err, ctx)
		return
	}

	for _, bill := range *bills{
		if bill.Frequency == models.ONCE_FREQUENCY || 
			bill.NextSequenceId != "" || 
			bill.DueDate.IsZero() || 
			time.Since(bill.DueDate) < 0 {

			continue
		}

		if time.Since(bill.DueDate) >= 0 {
			newBillId, err := s.CreateRecurringBill(ctx, &bill)
			if err != nil {
				log.Println("error while creating new bill id:", newBillId, "error:", err, ctx)
				continue
			}
			bill.NextSequenceId = newBillId
			_, err = s.DB.UpdateBillRecordById(ctx, bill.BillId, &bill)
			if err != nil {
				log.Println("error while updating budget id:", bill.BillId, "error:", err, ctx)
				continue
			}
			successfulCounter++
		}
	}

	log.Println("complete bill maintainer with successfully updating", successfulCounter, "bills", ctx)
}
