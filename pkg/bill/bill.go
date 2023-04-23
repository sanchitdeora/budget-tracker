package bill

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/utils"
)

type Service interface {
	GetBills(ctx context.Context, bill *[]models.Bill) (error)
	GetBillById(ctx context.Context, id string, bill *models.Bill) (error)
	CreateBillByUser(ctx context.Context, bill models.Bill) (string, error)
	CreateBill(ctx context.Context, bill models.Bill) (string, error)
	UpdateBillById(ctx context.Context, id string, bill models.Bill) (string, error)
	UpdateBillIsPaid(ctx context.Context, id string) (string, error)
	UpdateBillIsUnpaid(ctx context.Context, id string) (string, error)
	DeleteBillById(ctx context.Context, id string) (string, error)
}

type Opts struct {
	TransactionService transaction.Service
}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}


func (s *serviceImpl) GetBills(ctx context.Context, bill *[]models.Bill) (error) {
	// TODO: input validation
	return db.GetAllBillRecords(ctx, bill)
}

func (s *serviceImpl) GetBillById(ctx context.Context, id string, bill *models.Bill) (error) {
	// TODO: input validation
	return db.GetBillRecordById(ctx, id, bill)
}

func (s *serviceImpl) CreateBillByUser(ctx context.Context, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetByUser()
	bill.SetCategory()
	bill.SetFrequency()
	return db.InsertBillRecord(ctx, bill)
}

func (s *serviceImpl) CreateBill(ctx context.Context, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetCategory()
	bill.SetFrequency()
	return db.InsertBillRecord(ctx, bill)
}

func (s *serviceImpl) UpdateBillById(ctx context.Context, id string, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetCategory()
	bill.SetFrequency()
	return db.UpdateBillRecordById(ctx, id, bill)
}

func (s *serviceImpl) UpdateBillIsPaid(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	datePaid := time.Now()

	var bill models.Bill
	db.GetBillRecordById(ctx, id, &bill)
	
	if !bill.IsPaid {

		var newTransaction models.Transaction
		newTransaction.FromBill(bill, datePaid)
		_, err := transaction.Service.CreateTransaction(transaction.NewService(&transaction.Opts{}), ctx, newTransaction)
		if err != nil {
			log.Fatal(err)
		}
	}

	if bill.Frequency != models.ONCE_FREQUENCY {
		// create new bill entry for next frequency period
		newDueDate, err := utils.CalculateEndDateWithFrequency(bill.DueDate, bill.Frequency)
		if err != nil {
			log.Fatal(err)
		}

		newBill := bill
		newBill.BillId = ""
		newBill.DueDate = newDueDate
		newBill.IsPaid = false
		newBill.CreationTime = time.Now()
		newBill.SequenceNumber = bill.SequenceNumber + 1
		newBill.SequenceStartId = bill.SequenceStartId

		_, err = s.CreateBill(ctx, newBill)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db.UpdateBillRecordIsPaid(ctx, id, datePaid)
}

func (s *serviceImpl) UpdateBillIsUnpaid(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.UpdateBillRecordIsUnpaid(ctx, id)
}

func (s *serviceImpl) DeleteBillById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteBillRecordById(ctx, id)
}