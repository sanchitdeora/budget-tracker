package db

import (
	"context"

	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	budgetDatabase   	  *mongo.Database
	userCollection   	  *mongo.Collection
	surveyCollection	  *mongo.Collection
	transactionCollection *mongo.Collection
	billCollection *mongo.Collection
)

const (
	EMAIL_KEY = "email"
	TITLE_KEY = "title"
	CATEGORY_KEY = "category"
	AMOUNT_KEY = "amount"
	AMOUNT_DUE_KEY = "amount_due"
	DATE_KEY = "date"
	DUE_DATE_KEY = "due_date"
	DATE_PAID_KEY = "date_paid"
	FREQUENCY_KEY = "frequency"
	NOTE_KEY = "note"
	IS_PAID_KEY = "is_paid"
	ACCOUNT_KEY = "account"

	// Transaction constants
	TRANSACTION_ID_KEY = "transaction_id"
	TRANSACTION_PREFIX = "T-"
	TRANSACTION_TYPE_KEY = "type"
	
	// Bill constants
	BILL_ID_KEY = "bill_id"
	BILL_PREFIX = "B-"
)

func Init() (*mongo.Client, context.Context, error) {

	uri := "mongodb://localhost"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	budgetDatabase = client.Database("budget-tracker")

	userCollection = budgetDatabase.Collection("user_table")
	surveyCollection = budgetDatabase.Collection("survey_table")
	transactionCollection = budgetDatabase.Collection("trasanction_table")
	billCollection = budgetDatabase.Collection("bill_table")

	return client, ctx, err
}

func GetUserRecordByEmail(ctx context.Context, key string) ([]byte, error) {
	var user bson.M
	if err := userCollection.FindOne(ctx, bson.M{EMAIL_KEY: key}).Decode(&user); err != nil {
		log.Print("Error trying to get user from db", err)
		return nil, err
	}

	userJSON, err := bson.MarshalExtJSON(&user, true, true)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return userJSON, nil
}

func UpdateUserRecord(ctx context.Context, value string, update primitive.D) error {

	filter := bson.D{primitive.E{Key: EMAIL_KEY, Value: value}}

	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error trying to insert bdoc in mongo", err)
		return err
	}

	return nil
}

func AddUserRecord(ctx context.Context, record []byte) error {
	return addRecordToCollection(ctx, record, userCollection)
}

func AddSurveyRecord(ctx context.Context, record []byte) error {
	return addRecordToCollection(ctx, record, surveyCollection)
}

func addRecordToCollection(ctx context.Context, record []byte, collection *mongo.Collection) error {
	var bdoc interface{}
	err := bson.UnmarshalExtJSON(record, true, &bdoc)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := collection.InsertOne(ctx, bdoc)
	if err != nil {
		log.Println("Error trying to insert bdoc in mongo", err)
		return err
	}
	fmt.Println(result.InsertedID)

	return nil
}
