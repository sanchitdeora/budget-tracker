package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTransactions(ctx context.Context, transactions *[]models.Transaction) error {
	cur, err := transactionCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return err
	}

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(ctx)

	err = utils.ConvertBsonToStruct(results, transactions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Get All transaction. Count of elements: %v\n", len(results))
	return nil
	
}

func GetTransactionRecordById(ctx context.Context, key string, transaction *models.Transaction) error {
	var result bson.M

	filter := bson.M{TRANSACTION_ID_KEY: key}
	err := transactionCollection.FindOne(ctx, filter).Decode(&result)
	if len(result) == 0 {
		return nil
	}
	if err != nil {
		log.Println(err)
		return err
	}

	if err := utils.ConvertBsonToStruct(result, transaction); err != nil {
		log.Println(err)
		return err
	}

	return nil
	
}

func InsertTransactionRecord(ctx context.Context, transaction models.Transaction) (string, error) {
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
		log.Fatal(err)
	}
	fmt.Printf("Created transaction. ResultId: %v TransactionId: %v\n", result.InsertedID, transactionId)
	return transactionId, err
}

func UpdateTransactionRecordById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
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
		log.Fatal(err)
	}
	fmt.Printf("Updated transaction. ModifiedCount: %v TransactionId: %v\n", result.ModifiedCount, id)
	return id, err
}

func DeleteTransactionRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: TRANSACTION_ID_KEY, Value: id}}

	result, err := transactionCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted transaction. DeletedCount: %v TransactionId: %v\n", result.DeletedCount, id)
	return id, err
}