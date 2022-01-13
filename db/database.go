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

const userPrefix string = "U"
const surveyPrefix string = "S"

func AddUser(ctx context.Context, user models.User) error {
	user.UserID = userPrefix + uuid.NewString()

	record, err := bson.Marshal(user)
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

func GetUser(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	userJSON, err := GetUserRecord(ctx, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	userJSON, err := GetUserRecordByEmail(ctx, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func AddSurvey(ctx context.Context, survey *models.Survey) error {

	survey.SurveyID = surveyPrefix + uuid.NewString()
	log.Println("survey details", survey)
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

func GetSurvey(ctx context.Context, surveyID string) (*models.Survey, error) {
	var survey models.Survey
	surveyJSON, err := GetSurveyRecord(ctx, surveyID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("before survey", survey, err)
	err = json.Unmarshal(surveyJSON, &survey)
	log.Println("after survey", survey, err)
	if err != nil {
		log.Println(err)
		return &survey, err
	}
	return &survey, nil
}
