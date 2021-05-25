package repository

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

func Init() (*mongo.Client, context.Context, error){

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
	userCollection = budgetDatabase.Collection("account_table")

	return client, ctx, err
}

func CreateRecord(ctx context.Context, record []byte) error {
	var bdoc interface{}
	// fmt.Println("Unmarshal json to bson")
	err := bson.UnmarshalExtJSON(record, true, &bdoc)
	if err != nil {
		log.Fatal(err)
	}
	
	accountResult, err := userCollection.InsertOne(ctx, bdoc)
	if err != nil {
		log.Fatal("Error trying to insert bdoc in mongo", err)
	}
	fmt.Println(accountResult.InsertedID)
	
	return nil
}

func GetRecord(ctx context.Context, key string) ([]byte, error) {
	var user bson.M
	if err := userCollection.FindOne(ctx, bson.M{"email": key}).Decode(&user); err != nil {
		log.Fatal("Error trying to get user from db", err)
	}
	
	userJSON, err := bson.MarshalExtJSON(&user, true, true)
	if err != nil {
		log.Fatal(err)
	}
	return userJSON, nil
}
