package db

import (
	"context"

	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var budgetDatabase *mongo.Database
var userCollection *mongo.Collection
var surveyCollection *mongo.Collection

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

	// fmt.Println("Start creating database")
	budgetDatabase = client.Database("budget-tracker")

	userCollection = budgetDatabase.Collection("user_table")
	surveyCollection = budgetDatabase.Collection("survey_table")

	return client, ctx, err
}

func GetUserRecord(ctx context.Context, key string) ([]byte, error) {
	return getRecordFromCollection(ctx, key, userCollection)
}

func AddUserRecord(ctx context.Context, record []byte) error {
	return addRecordToCollection(ctx, record, userCollection)
}

func AddSurveyRecord(ctx context.Context, record []byte) error {
	return addRecordToCollection(ctx, record, surveyCollection)
}

func getRecordFromCollection(ctx context.Context, key string, collection *mongo.Collection) ([]byte, error) {
	var user bson.M
	if err := collection.FindOne(ctx, bson.M{"email": key}).Decode(&user); err != nil {
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
