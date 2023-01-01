package bill

import (
	"context"
	"errors"
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
		newDueDate, err := calculateNewBillDate(bill.DueDate, bill.Frequency)
		if err != nil {
			log.Fatal(err)
		}

		newBill := bill
		newBill.BillId = ""
		newBill.DueDate = newDueDate
		newBill.IsPaid = false
		
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

func calculateNewBillDate(currDate time.Time, freq string) (time.Time, error) {
	err := errors.New("provided Frequency is not found in the frequency map")
	
	if !utils.Contains(models.BillFrequencyMap, freq) || models.ONCE_FREQUENCY == freq {
		return currDate, err
	}
	if freq == models.DAILY_FREQUENCY {
		return currDate.AddDate(0, 0, 1), nil
	} else if freq == models.WEEKLY_FREQUENCY {
		return currDate.AddDate(0, 0, 7), nil
	} else if freq == models.BI_WEEKLY_FREQUENCY {
		return currDate.AddDate(0, 0, 14), nil
	} else if freq == models.MONTHLY_FREQUENCY {
		return currDate.AddDate(0, 1, 0), nil
	} else if freq == models.BI_MONTHLY_FREQUENCY {
		return currDate.AddDate(0, 2, 0), nil
	} else if freq == models.QUATERLY_FREQUENCY {
		return currDate.AddDate(0, 3, 0), nil
	} else if freq == models.HALF_YEARLY_FREQUENCY {
		return currDate.AddDate(0, 6, 0), nil
	} else if freq == models.YEARLY_FREQUENCY {
		return currDate.AddDate(1, 0, 0), nil
	} else {
		return currDate, err
	}

}