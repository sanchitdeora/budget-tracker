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
var accountsTable *mongo.Collection

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
	accountsTable = budgetDatabase.Collection("account_table")

	return client, ctx, err
}

func CreateAccountRecord(ctx context.Context, record []byte) {
	var bdoc interface{}
	// fmt.Println("Unmarshal json to bson")
	err := bson.UnmarshalExtJSON(record, true, &bdoc)
	if err != nil {
		log.Fatal(err)
	}
	accountResult, err := accountsTable.InsertOne(ctx, bdoc)
	if err != nil {
		log.Fatal("Error trying to insert bdoc in mongo", err)
	}
	fmt.Println(accountResult.InsertedID)
}
