package repository

import (
	"context"
	"encoding/json"

	"log"

	"github.com/sanchitdeora/budget-tracker/models"
)

func CreateAccountRecords(ctx context.Context, user models.User) {
	record, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}		
	
	// fmt.Println("Inside db create account record", string(record))

	CreateAccountRecord(ctx, record)
}