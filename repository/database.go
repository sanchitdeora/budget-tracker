package repository

import (
	"context"
	"encoding/json"
	"log"
	"github.com/sanchitdeora/budget-tracker/models"
)

func CreateUserRecord(ctx context.Context, user models.User) error {
	record, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}		
	
	// fmt.Println("Inside db create account record", string(record))
	err = CreateRecord(ctx, record)
	if err != nil {
		log.Fatal(err)
	}		
	return err
}

func GetLoginInfo(ctx context.Context, login models.Login) (models.Login, error) {
	userJSON, err := GetRecord(ctx, login.Email)
	if err != nil {
		log.Fatal(err)
	}

	var loginDB models.Login 
	err = json.Unmarshal(userJSON, &loginDB)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("After unmarshall", loginDB)
	return loginDB, nil
}