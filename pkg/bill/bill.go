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
	CreateBill(ctx context.Context, bill *models.Bill) (string, error)
	UpdateBillById(ctx context.Context, id string, bill *models.Bill) (string, error)
	UpdateBillIsPaid(ctx context.Context, id string) (string, error)
	UpdateBillIsUnpaid(ctx context.Context, id string) (string, error)
	DeleteBillById(ctx context.Context, id string) (string, error)
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

func (s *serviceImpl) CreateBill(ctx context.Context, bill *models.Bill) (string, error) {
	if !bill.IsValid() {
		return "", exceptions.ErrValidationError
	}

	bill.SetCategory()
	bill.SetFrequency()
	return s.DB.InsertBillRecord(ctx, bill)
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
	bill, err := s.DB.GetBillRecordById(ctx, id)
	if err != nil {
		log.Println("error while fetching bill record for billId:", id, "error:", err, ctx)
		return "", err
	}

	nextBillId := bill.NextSequenceId
	if bill.NextSequenceId == "" {
		datePaid := time.Now()
		if !bill.IsPaid {
			var newTransaction models.Transaction
			newTransaction.FromBill(*bill, datePaid)
			_, err := s.TransactionService.CreateTransaction(ctx, &newTransaction)
			if err != nil {
				log.Println("error while creating transaction record", newTransaction, "error:", err, ctx)
				return "", err
			}
		}
		
		if bill.Frequency != models.ONCE_FREQUENCY {
			// create new bill entry for next frequency period
			bill.SetFrequency()
			newDueDate:= utils.CalculateEndDateWithFrequency(bill.DueDate, bill.Frequency)
			
			newBill := bill
			newBill.BillId = ""
			newBill.DueDate = newDueDate
			newBill.IsPaid = false
			newBill.CreationTime = time.Now()
			newBill.SequenceNumber = bill.SequenceNumber + 1
			newBill.SequenceStartId = bill.SequenceStartId

			nextBillId, err = s.CreateBill(ctx, newBill)
			if err != nil {
				log.Println("error while creating new bill record:", newBill, "error:", err, ctx)
				return "", err
			}
		}
	}

	return s.DB.UpdateBillRecordIsPaid(ctx, id, nextBillId, datePaid)
}

func (s *serviceImpl) UpdateBillIsUnpaid(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return "", exceptions.ErrValidationError
	}

	return s.DB.UpdateBillRecordIsUnpaid(ctx, id)
}

func (s *serviceImpl) DeleteBillById(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("missing Bill Id", ctx)
		return "", exceptions.ErrValidationError
	}

	return s.DB.DeleteBillRecordById(ctx, id)
}