package bill

import (
	"context"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
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
	bill.SetHowOften()
	return db.InsertBillRecord(ctx, bill)
}

func updateBillById(ctx context.Context, id string, bill models.Bill) (string, error) {
	// TODO: input validation
	bill.SetCategory()
	bill.SetHowOften()
	return db.UpdateBillRecordById(ctx, id, bill)
}

func deleteBillById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteBillRecordById(ctx, id)
}
