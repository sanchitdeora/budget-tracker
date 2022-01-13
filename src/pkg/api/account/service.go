package account

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func GetAccount(ctx context.Context, userID string) (*models.AccountResponse, error) {
	user, err := db.GetUser(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid user request")
	}
	var userRes models.UserResponse
	userRes.UserID = userID
	userRes.Firstname = user.Firstname
	userRes.Lastname = user.Lastname
	userRes.Email = user.Email
	userRes.Password = user.Password
	userRes.PhoneNumber = user.PhoneNumber
	userRes.DateOfBirth = user.DateOfBirth

	log.Println(userRes)

	if !user.SurveyComplete {
		return &models.AccountResponse{User: userRes}, nil
	}
	survey, err := db.GetSurvey(ctx, user.SurveyID)
	log.Println("survey", survey)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if survey == nil {
		return nil, errors.New("invalid survey request")
	}
	var surveyRes models.SurveyResponse
	surveyRes.SavingsType = survey.SavingsType
	surveyRes.MonthlyLimit = survey.MonthlyLimit
	surveyRes.MonthlyIncome = survey.MonthlyIncome

	log.Println(userRes)

	return &models.AccountResponse{User: userRes, Survey: surveyRes}, nil
}

func LoginService(ctx context.Context, login *models.Login) (*models.LoginResponse, error) {
	// improve validation logic

	user, err := db.GetUserByEmail(ctx, login.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if user.Password != login.Password {
		return nil, errors.New("invaid login")
	}

	var resp models.LoginResponse

	resp.Message = "login response"
	resp.Token = tokenGenerator()
	resp.Email = login.Email
	resp.Name = user.Lastname + ", " + user.Firstname
	resp.SurveyComplete = user.SurveyComplete

	return &resp, err
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
