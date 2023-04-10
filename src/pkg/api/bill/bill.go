package bill

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/transaction"
	"github.com/sanchitdeora/budget-tracker/src/utils"
)

func getBills(ctx context.Context, bill *[]models.Bill) (error) {	
	// TODO: input validation
	return db.GetAllBillRecords(ctx, bill)
}

func getBillById(ctx context.Context, id string, bill *models.Bill) (error) {
	// TODO: input validation
	return db.GetBillRecordById(ctx, id, bill)
}

func createBillByUser(ctx context.Context, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetByUser()
	bill.SetCategory()
	bill.SetFrequency()
	return db.InsertBillRecord(ctx, bill)
}

func createBill(ctx context.Context, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetCategory()
	bill.SetFrequency()
	return db.InsertBillRecord(ctx, bill)
}

func updateBillById(ctx context.Context, id string, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetCategory()
	bill.SetFrequency()
	return db.UpdateBillRecordById(ctx, id, bill)
}

func updateBillIsPaid(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	datePaid := time.Now()

	var bill models.Bill
	db.GetBillRecordById(ctx, id, &bill)
	
	if !bill.IsPaid {

		var newTransaction models.Transaction
		newTransaction.FromBill(bill, datePaid)
		_, err := transaction.CreateTransactionRecord(ctx, newTransaction)
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
		
		_, err = createBill(ctx, newBill)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db.UpdateBillRecordIsPaid(ctx, id, datePaid)
}

func updateBillIsUnpaid(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.UpdateBillRecordIsUnpaid(ctx, id)
}

func deleteBillById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteBillRecordById(ctx, id)
}