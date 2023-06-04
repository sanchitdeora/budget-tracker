package db

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER_PREFIX string = "U-"
const SURVEY_PREFIX string = "S-"

const (
	ACCOUNT_KEY = "account"
	AMOUNT_KEY = "amount"
	AMOUNT_DUE_KEY = "amount_due"
	CATEGORY_KEY = "category"
	CREATION_TIME_KEY = "creation_time"
	CURRENT_AMOUNT_KEY = "current_amount"
	DATE_KEY = "date"
	DATE_PAID_KEY = "date_paid"
	DUE_DATE_KEY = "due_date"
	EMAIL_KEY = "email"
	EXPIRATION_TIME_KEY = "expiration_time"
	FREQUENCY_KEY = "frequency"
	INCOME_KEY = "income"
	IS_PAID_KEY = "is_paid"
	NAME_KEY = "name"
	NOTE_KEY = "note"
	SAVINGS_KEY = "savings"
	SEQUENCE_NUMBER_KEY = "sequence_no"
	SEQUENCE_START_ID_KEY = "sequence_start_id"
	TARGET_AMOUNT_KEY = "target_amount"
	TARGET_DATE_KEY = "target_date"
	TITLE_KEY = "title"

	// Transaction constants
	TRANSACTION_ID_KEY = "transaction_id"
	TRANSACTION_PREFIX = "T-"
	TRANSACTION_TYPE_KEY = "type"
	
	// Bill constants
	BILL_ID_KEY = "bill_id"
	BILL_PREFIX = "B-"


	// Budget constants
	BUDGET_ID_KEY = "budget_id"
	BUDGET_NAME_KEY = "budget_name"
	BUDGET_PREFIX = "BG-"
	BUDGET_INCOME_MAP_KEY = "income_map"
	BUDGET_EXPENSE_MAP_KEY = "expense_map"
	BUDGET_GOAL_MAP_KEY = "goal_map"
	
	// Goal constants
	GOAL_ID_KEY = "goal_id"
	GOAL_NAME_KEY = "goal_name"
	GOAL_PREFIX = "G-"
	BUDGET_ID_LIST_KEY = "budget_id_list"
)

type Database interface {
	// transaction
	GetAllTransactionRecords(ctx context.Context, transactions *[]models.Transaction) error
	GetTransactionRecordById(ctx context.Context, key string, transaction *models.Transaction) error
	InsertTransactionRecord(ctx context.Context, transaction models.Transaction) (string, error)
	UpdateTransactionRecordById(ctx context.Context, id string, transaction models.Transaction) (string, error)
	DeleteTransactionRecordById(ctx context.Context, id string) (string, error)

	// bill
	GetAllBillRecords(ctx context.Context, bills *[]models.Bill) error
	GetBillRecordById(ctx context.Context, id string, bill *models.Bill) error
	InsertBillRecord(ctx context.Context, bill models.Bill) (string, error)
	UpdateBillRecordById(ctx context.Context, id string, bill models.Bill) (string, error)
	UpdateBillRecordIsPaid(ctx context.Context, id string, datePaid time.Time) (string, error)
	UpdateBillRecordIsUnpaid(ctx context.Context, id string) (string, error)
	DeleteBillRecordById(ctx context.Context, id string) (string, error)

	// budget
	GetAllBudgetRecords(ctx context.Context, budgets *[]models.Budget) error
	GetBudgetRecordById(ctx context.Context, key string) (*models.Budget, error)
	InsertBudgetRecord(ctx context.Context, budget models.Budget) (string, error)
	UpdateBudgetRecordById(ctx context.Context, id string, budget models.Budget) (string, error)
	DeleteBudgetRecordById(ctx context.Context, id string) (string, error)

	// goal
	GetAllGoalRecords(ctx context.Context, goals *[]models.Goal) error
	GetGoalRecordById(ctx context.Context, key string) (*models.Goal, error)
	InsertGoalRecord(ctx context.Context, goal models.Goal) (string, error)
	UpdateGoalRecordById(ctx context.Context, id string, goal models.Goal) (string, error)
	DeleteGoalRecordById(ctx context.Context, id string) (string, error)
}

type DatabaseImpl struct {}

func NewDatabase() Database {
	return &DatabaseImpl{}
}

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
