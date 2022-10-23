package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/src/models"
	"github.com/sanchitdeora/budget-tracker/src/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllBillRecords(ctx context.Context, bills *[]models.Bill) error {
	cur, err := billCollection.Find(ctx, bson.D{})
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

	err = utils.ConvertBsonToStruct(results, bills)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Get All bills. Count of elements: %v\n", len(results))
	return nil
	
}

func GetBillRecordById(ctx context.Context, id string, bill *models.Bill) error {
	filter := bson.M{billIdKey: id}
	if err := billCollection.FindOne(ctx, filter).Decode(&bill); err != nil {
		log.Println(err)
		return err
	}

	var result bson.M
	if err := utils.ConvertBsonToStruct(result, bill); err != nil {
		log.Println(err)
		return err
	}

	return nil
	
}

func InsertBillRecord(ctx context.Context, bill models.Bill) (string, error) {
	billId := billPrefix + uuid.NewString()
	data := bson.D{
		{Key: billIdKey, Value: billId},
		{Key: titleKey, Value: bill.Title},
		{Key: categoryKey, Value: bill.Category},
		{Key: amountDueKey, Value: bill.AmountDue},
		{Key: dueDataKey, Value: bill.DueDate},
		{Key: noteKey, Value: bill.Note},
		{Key: isPaidKey, Value: bill.IsPaid},
	}

	result, err := billCollection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created bill. ResultId: %v BillId: %v\n", result.InsertedID, billId)
	return billId, err
}

func UpdateBillRecordById(ctx context.Context, id string, bill models.Bill) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: titleKey, Value: bill.Title},
			{Key: categoryKey, Value: bill.Category},
			{Key: amountDueKey, Value: bill.AmountDue},
			{Key: dueDataKey, Value: bill.DueDate},
			{Key: noteKey, Value: bill.Note},
			{Key: isPaidKey, Value: bill.IsPaid},
		}},
	}
	filter := bson.D{{Key: billIdKey, Value: id}}

	result, err := billCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated bill. ModifiedCount: %v BillId: %v\n", result.ModifiedCount, id)
	return id, err
}

func DeleteBillRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: billIdKey, Value: id}}

	result, err := billCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted bill. DeletedCount: %v BillId: %v\n", result.DeletedCount, id)
	return id, err
}