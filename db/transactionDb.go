package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	transactionIdKey = "transactionId"
	transactionPrefix = "T-"
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

	records, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(records, &transactions)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Get All transaction. Count of elements: %v\n", len(results))
	return nil
	
}

func GetTransactionRecordById(ctx context.Context, key string, transaction *models.Transaction) error {
	filter := bson.M{transactionIdKey: key}
	if err := transactionCollection.FindOne(ctx, filter).Decode(&transaction); err != nil {
		log.Println(err)
		return err
	}

	var result bson.M
	transactionJson, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(transactionJson, &transaction)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
	
}

func InsertTransactionRecord(ctx context.Context, transaction models.Transaction) (string, error) {
	transactionId := transactionPrefix + uuid.NewString()
	data := bson.D{
		{Key: "transactionId", Value: transactionId},
		{Key: "title", Value: transaction.Title},
		{Key: "category", Value: transaction.Category},
		{Key: "amount", Value: transaction.Amount},
		{Key: "date", Value: transaction.Date},
		{Key: "description", Value: transaction.Description},
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
			{Key: "transactionId", Value: id},
			{Key: "title", Value: transaction.Title},
			{Key: "category", Value: transaction.Category},
			{Key: "amount", Value: transaction.Amount},
			{Key: "date", Value: transaction.Date},
			{Key: "description", Value: transaction.Description},
		}},
	}
	filter := bson.D{{Key: transactionIdKey, Value: id}}

	result, err := transactionCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated transaction. ModifiedCount: %v TransactionId: %v\n", result.ModifiedCount, id)
	return id, err
}

func DeleteTransactionRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: transactionIdKey, Value: id}}

	result, err := transactionCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted transaction. DeletedCount: %v TransactionId: %v\n", result.DeletedCount, id)
	return id, err
}