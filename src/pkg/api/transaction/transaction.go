package transaction

import (
	"context"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func getTransactions(ctx context.Context, transactions *[]models.Transaction) (error) {
	// TODO: input validation
	return db.GetAllTransactions(ctx, transactions)
}

func getTransactionById(ctx context.Context, id string, transaction *models.Transaction) (error) {
	// TODO: input validation
	return db.GetTransactionRecordById(ctx, id, transaction)
}

func CreateTransactionRecord(ctx context.Context, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return db.InsertTransactionRecord(ctx, transaction)
}

func updateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return db.UpdateTransactionRecordById(ctx, id, transaction)
}

func deleteTransactionById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteTransactionRecordById(ctx, id)
}
