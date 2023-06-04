package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DatabaseImpl) GetAllBudgetRecords(ctx context.Context, budgets *[]models.Budget) error {
	cur, err := budgetCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return err
	}

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(ctx)

	err = utils.ConvertBsonToStruct(results, budgets)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Get All budget. Count of elements: %v\n", len(results))
	return nil
	
}

func (db *DatabaseImpl) GetBudgetRecordById(ctx context.Context, key string) (*models.Budget, error) {
	var result bson.M
	var budget models.Budget

	filter := bson.M{BUDGET_ID_KEY: key}
	err := budgetCollection.FindOne(ctx, filter).Decode(&result)
	if len(result) == 0 {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := utils.ConvertBsonToStruct(result, &budget); err != nil {
		log.Println(err)
		return nil, err
	}

	return &budget, nil
	
}

func (db *DatabaseImpl) InsertBudgetRecord(ctx context.Context, budget models.Budget) (string, error) {
	if budget.BudgetId == "" {
		budget.BudgetId = BUDGET_PREFIX + uuid.NewString()
	}	
	data := bson.D{
		{Key: BUDGET_ID_KEY, Value: budget.BudgetId},
		{Key: BUDGET_NAME_KEY, Value: budget.BudgetName},
		{Key: BUDGET_INCOME_MAP_KEY, Value: budget.IncomeMap},
		{Key: BUDGET_EXPENSE_MAP_KEY, Value: budget.ExpenseMap},
		{Key: BUDGET_GOAL_MAP_KEY, Value: budget.GoalMap},
		{Key: FREQUENCY_KEY, Value: budget.Frequency},
		{Key: SAVINGS_KEY, Value: budget.Savings},
		{Key: CREATION_TIME_KEY, Value: budget.CreationTime},
		{Key: EXPIRATION_TIME_KEY, Value: budget.ExpirationTime},
		{Key: SEQUENCE_NUMBER_KEY, Value: budget.SequenceNumber},
		{Key: SEQUENCE_START_ID_KEY, Value: budget.SequenceStartId},
	}

	result, err := budgetCollection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created budget. ResultId: %v BudgetId: %v\n", result.InsertedID, budget.BudgetId)
	return budget.BudgetId, err
}

func (db *DatabaseImpl) UpdateBudgetRecordById(ctx context.Context, id string, budget models.Budget) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: BUDGET_NAME_KEY, Value: budget.BudgetName},
			{Key: BUDGET_INCOME_MAP_KEY, Value: budget.IncomeMap},
			{Key: BUDGET_EXPENSE_MAP_KEY, Value: budget.ExpenseMap},
			{Key: BUDGET_GOAL_MAP_KEY, Value: budget.GoalMap},
			{Key: FREQUENCY_KEY, Value: budget.Frequency},
			{Key: SAVINGS_KEY, Value: budget.Savings},
		}},
	}
	filter := bson.D{{Key: BUDGET_ID_KEY, Value: id}}

	result, err := budgetCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated budget. ModifiedCount: %v BudgetId: %v\n", result.ModifiedCount, id)
	return id, err
}

func (db *DatabaseImpl) DeleteBudgetRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: BUDGET_ID_KEY, Value: id}}

	result, err := budgetCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted budget. DeletedCount: %v BudgetId: %v\n", result.DeletedCount, id)
	return id, err
}