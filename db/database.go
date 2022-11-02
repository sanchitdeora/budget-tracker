package db

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER_PREFIX string = "U-"
const SURVEY_PREFIX string = "S-"

func AddUser(ctx context.Context, user models.User) error {
	user.UserID = USER_PREFIX + uuid.NewString()

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

func UpdateUser(ctx context.Context, user models.User) error {

	update := bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{primitive.E{
			Key:   "surveyId",
			Value: user.SurveyID,
		}, primitive.E{
			Key:   "surveyComplete",
			Value: user.SurveyComplete,
		}},
	}}

	err := UpdateUserRecord(ctx, user.Email, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	userRecord, err := GetUserRecordByEmail(ctx, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(userRecord, &user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func AddSurvey(ctx context.Context, survey *models.Survey) error {

	survey.SurveyID = SURVEY_PREFIX + uuid.NewString()
	record, err := json.Marshal(survey)
	if err != nil {
		log.Println(err)
		return err
	}

	err = AddSurveyRecord(ctx, record)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetLoginInfo(ctx context.Context, login *models.Login) (*models.Login, error) {
	var loginDB models.Login
	userRecord, err := GetUserRecordByEmail(ctx, login.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(userRecord, &loginDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// fmt.Println("After unmarshall", loginDB)
	return &loginDB, nil
}
