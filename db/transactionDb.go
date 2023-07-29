package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/sanchitdeora/budget-tracker/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DatabaseImpl) GetAllTransactionRecords(ctx context.Context) (*[]models.Transaction, error) {
	cur, err := transactionCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// fmt.Println("Start getting all transactions")

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println("error while fetching all transactions, error: ", err)
			return nil, err
		}
		results = append(results, result)
	}

	// fmt.Println("Got Results for all transactions in bson: ", len(results))

	if err := cur.Err(); err != nil {
		log.Println("error while fetching all transactions, error: ", err)
		return nil, err
	}
	cur.Close(ctx)

	var transactions []models.Transaction
	if err := utils.ConvertBsonToStruct(results, &transactions); err != nil {
		log.Println("error while converting bson to struct, error: ", err)
		return nil, err
	}

	log.Printf("Get All transaction. Count of elements: %v\n", len(results))
	return &transactions, nil
}

func (db *DatabaseImpl) GetTransactionRecordById(ctx context.Context, key string) (*models.Transaction, error) {
	var result bson.M

	filter := bson.M{TRANSACTION_ID_KEY: key}
	err := transactionCollection.FindOne(ctx, filter).Decode(&result)
	if len(result) == 0 {
		log.Println("transaction not found for id: ", key)
		return nil, exceptions.ErrTransactionNotFound
	}
	if err != nil {
		log.Println("error while fetching transaction by id: ", key, " error: ", err)
		return nil, err
	}

	var transaction models.Transaction
	if err := utils.ConvertBsonToStruct(result, &transaction); err != nil {
		log.Println("error while converting bson to struct, error: ", err)
		return nil, err
	}

	return &transaction, nil
}

func (db *DatabaseImpl) GetAllTransactionRecordsByDateRange(ctx context.Context, startDate time.Time, endDate time.Time) (*[]models.Transaction, error) {
	filter := bson.M{
        "date": bson.M{
            "$gt": startDate,
            "$lt": endDate,
        },
	}
	
	// fmt.Println("filter here: ", filter)
	cur, err := transactionCollection.Find(ctx, filter)
	if err != nil {
		log.Println("error while fetching transaction by date, startDate: ", startDate, "endDate: ", endDate, " error: ", err)
		return nil, err
	}
	
	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println("error while fetching transactions by date, error: ", err)
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Println("error while fetching transactions by date, error: ", err)
		return nil, err
	}
	cur.Close(ctx)
	
	var transactions []models.Transaction
	if err := utils.ConvertBsonToStruct(results, &transactions); err != nil {
		log.Println("error while converting bson to struct, error: ", err)
		return nil, err
	}

	fmt.Printf("Get All transaction by date. Count of elements: %v\n", len(results))
	return &transactions, nil

}

func (db *DatabaseImpl) InsertTransactionRecord(ctx context.Context, transaction models.Transaction) (string, error) {
	transactionId := TRANSACTION_PREFIX + uuid.NewString()
	data := bson.D{
		{Key: TRANSACTION_ID_KEY, Value: transactionId},
		{Key: TITLE_KEY, Value: transaction.Title},
		{Key: CATEGORY_KEY, Value: transaction.Category},
		{Key: AMOUNT_KEY, Value: transaction.Amount},
		{Key: DATE_KEY, Value: transaction.Date},
		{Key: TRANSACTION_TYPE_KEY, Value: transaction.Type},
		{Key: ACCOUNT_KEY, Value: transaction.Account},
		{Key: NOTE_KEY, Value: transaction.Note},
	}

	result, err := transactionCollection.InsertOne(ctx, data)
	if err != nil {
		log.Println("error while inserting transaction, error: ", err)
		return "", err
	}
	fmt.Printf("Created transaction. ResultId: %v TransactionId: %v\n", result.InsertedID, transactionId)
	return transactionId, err
}

func (db *DatabaseImpl) UpdateTransactionRecordById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: TITLE_KEY, Value: transaction.Title},
			{Key: CATEGORY_KEY, Value: transaction.Category},
			{Key: AMOUNT_KEY, Value: transaction.Amount},
			{Key: DATE_KEY, Value: transaction.Date},
			{Key: TRANSACTION_TYPE_KEY, Value: transaction.Type},
			{Key: ACCOUNT_KEY, Value: transaction.Account},
			{Key: NOTE_KEY, Value: transaction.Note},
		}},
	}
	filter := bson.D{{Key: TRANSACTION_ID_KEY, Value: id}}

	result, err := transactionCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Println("error while updating transaction with id: ", id, ", error: ", err)
		return "", err
	}
	fmt.Printf("Updated transaction. ModifiedCount: %v TransactionId: %v\n", result.ModifiedCount, id)
	return id, err
}

func (db *DatabaseImpl) DeleteTransactionRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: TRANSACTION_ID_KEY, Value: id}}

	result, err := transactionCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("error while deleting transaction with id: ", id, ", error: ", err)
		return "", err
	}

	fmt.Printf("Deleted transaction. DeletedCount: %v TransactionId: %v\n", result.DeletedCount, id)
	return id, err
}