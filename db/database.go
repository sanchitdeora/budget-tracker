package db

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

const userPrefix string = "U-"
const surveyPrefix string = "S-"

func AddUser(ctx context.Context, user models.User) error {
	user.UserID = userPrefix + uuid.NewString()

	record, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return err
	}

	// fmt.Println("Inside db create account record", string(record))
	err = AddUserRecord(ctx, record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func AddSurvey(ctx context.Context, survey models.Survey) error {
	survey.SurveyID = surveyPrefix + uuid.NewString()
	
	record, err := json.Marshal(survey)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Inside db create account record", string(record))
	err = AddSurveyRecord(ctx, record)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetLoginInfo(ctx context.Context, login models.Login) (*models.Login, error) {
	var loginDB models.Login
	userJSON, err := GetUserRecord(ctx, login.Email)
	if err != nil {
		log.Println(err)
		return &loginDB, err
	}

	err = json.Unmarshal(userJSON, &loginDB)
	if err != nil {
		log.Println(err)
		return &loginDB, err
	}
	// fmt.Println("After unmarshall", loginDB)
	return &loginDB, nil
}
