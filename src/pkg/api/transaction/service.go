package transaction

import (
	"context"
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func getTransactions(ctx context.Context) ([]models.Transaction, error) {
	var results []models.Transaction
	
	err := db.GetAllTransactions(ctx, &results)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}

func getTransactionById(ctx context.Context, id string) (models.Transaction, error) {
	// TODO: input validation
	var result models.Transaction
	
	err := db.GetTransactionRecordById(ctx, id, &result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}

func createTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return db.InsertTransactionRecord(ctx, transaction)
}

func updateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
	// TODO: input validation
	return db.UpdateTransactionRecordById(ctx, id, transaction)
}

func deleteTransactionById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteTransactionRecordById(ctx, id)
}
