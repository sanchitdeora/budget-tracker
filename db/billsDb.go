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

func (db *DatabaseImpl) GetAllBillRecords(ctx context.Context) (*[]models.Bill, error) {
	cur, err := billCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println("error while fetching all bills, error: ", err)
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Println("error while fetching all bills, error: ", err)
		return nil, err
	}
	cur.Close(ctx)

	var bills []models.Bill
	err = utils.ConvertBsonToStruct(results, &bills)
	if err != nil {
		log.Println("error while converting bson to struct, error: ", err)
		return nil, err
	}

	fmt.Printf("Get All bills. Count of elements: %v\n", len(results))
	return &bills, nil
	
}

func (db *DatabaseImpl) GetBillRecordById(ctx context.Context, key string) ( *models.Bill, error) {
	var result bson.M

	filter := bson.D{{Key: BILL_ID_KEY, Value: key}}
	err := billCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println("error while fetching bill by id: ", key, " error: ", err)
		return nil, err
	}
	if len(result) == 0 {
		log.Println("bill not found for id: ", key)
		return nil, exceptions.ErrBillNotFound
	}
	
	var bill models.Bill
	if err := utils.ConvertBsonToStruct(result, &bill); err != nil {
		log.Println("error while converting bson to struct, error: ", err)
		return nil, err
	}
	
	return &bill, nil
	
}

func (db *DatabaseImpl) InsertBillRecord(ctx context.Context, bill *models.Bill) (string, error) {
	billId := BILL_PREFIX + uuid.NewString()

	if bill.SequenceStartId == "" && bill.SequenceNumber == 0 {
		bill.SequenceStartId = billId
	}

	data := bson.D{
		{Key: BILL_ID_KEY, Value: billId},
		{Key: TITLE_KEY, Value: bill.Title},
		{Key: CATEGORY_KEY, Value: bill.Category},
		{Key: AMOUNT_DUE_KEY, Value: bill.AmountDue},
		{Key: DUE_DATE_KEY, Value: bill.DueDate},
		{Key: FREQUENCY_KEY, Value: bill.Frequency},
		{Key: NOTE_KEY, Value: bill.Note},
		{Key: IS_PAID_KEY, Value: bill.IsPaid},
		{Key: CREATION_TIME_KEY, Value: bill.CreationTime},
		{Key: SEQUENCE_START_ID_KEY, Value: bill.SequenceStartId},
		{Key: SEQUENCE_NUMBER_KEY, Value: bill.SequenceNumber},
	}

	result, err := billCollection.InsertOne(ctx, data)
	if err != nil {
		log.Println("error while inserting bill to unpaid, error: ", err)
		return "", err
	}
	fmt.Printf("Created bill. ResultId: %v BillId: %v\n", result.InsertedID, billId)
	return billId, err
}

func (db *DatabaseImpl) UpdateBillRecordById(ctx context.Context, id string, bill *models.Bill) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: TITLE_KEY, Value: bill.Title},
			{Key: CATEGORY_KEY, Value: bill.Category},
			{Key: AMOUNT_DUE_KEY, Value: bill.AmountDue},
			{Key: DUE_DATE_KEY, Value: bill.DueDate},
			{Key: FREQUENCY_KEY, Value: bill.Frequency},
			{Key: NOTE_KEY, Value: bill.Note},
			{Key: IS_PAID_KEY, Value: bill.IsPaid},
			{Key: CREATION_TIME_KEY, Value: bill.CreationTime},
			{Key: SEQUENCE_START_ID_KEY, Value: bill.SequenceStartId},
			{Key: SEQUENCE_NUMBER_KEY, Value: bill.SequenceNumber},
		}},
	}
	filter := bson.D{{Key: BILL_ID_KEY, Value: id}}

	result, err := billCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Println("error while updating bill with BillId: ", id, ", error: ", err)
		return "", err
	}
	fmt.Printf("Updated bill. ModifiedCount: %v BillId: %v\n", result.ModifiedCount, id)
	return id, err
}

func (db *DatabaseImpl) UpdateBillRecordIsPaid(ctx context.Context, id string, datePaid time.Time) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: IS_PAID_KEY, Value: true},
			{Key: DATE_PAID_KEY, Value: datePaid},
		}},
	}
	filter := bson.D{{Key: BILL_ID_KEY, Value: id}}

	result, err := billCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Println("error while updating bill to paid with BillId: ", id, ", error: ", err)
		return "", err
	}
	fmt.Printf("IsPaid set to true for bill. ModifiedCount: %v BillId: %v\n", result.ModifiedCount, id)
	return id, err
}

func (db *DatabaseImpl) UpdateBillRecordIsUnpaid(ctx context.Context, id string) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: IS_PAID_KEY, Value: false},
		}},
	}
	filter := bson.D{{Key: BILL_ID_KEY, Value: id}}

	result, err := billCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Println("error while updating bill to unpaid with BillId: ", id, ", error: ", err)
		return "", err
	}
	fmt.Printf("IsPaid set to false for bill. ModifiedCount: %v BillId: %v\n", result.ModifiedCount, id)
	return id, err
}

func (db *DatabaseImpl) DeleteBillRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: BILL_ID_KEY, Value: id}}

	result, err := billCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("error while deleting bill with BillId: ", id, ", error: ", err)
		return "", err
	}

	fmt.Printf("Deleted bill. DeletedCount: %v BillId: %v\n", result.DeletedCount, id)
	return id, err
}