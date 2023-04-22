package budget

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/goal"
	"github.com/sanchitdeora/budget-tracker/src/utils"
)

func GetBudgets(ctx context.Context, budget *[]models.Budget) (error) {
	// TODO: input validation
	return db.GetAllBudgets(ctx, budget)
}

func GetBudget(ctx context.Context, id string) (*models.Budget, error) {
	// TODO: input validation
	return db.GetBudgetRecordById(ctx, id)
}

func GetGoalMap(ctx context.Context, id string) ([]models.BudgetInputMap, error) {
	budget, err := GetBudget(ctx, id)
	return budget.GoalMap, err
}

func createBudgetByUser(ctx *gin.Context, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	budget.SetByUser()
	return db.InsertBudgetRecord(ctx, budget)
}

func createBudget(ctx *gin.Context, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	budget.AutoSet()
	return db.InsertBudgetRecord(ctx, budget)
}

func UpdateBudgetById(ctx *gin.Context, id string, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	
	currentGoalMap, err := GetGoalMap(ctx, id)
	if err != nil {
		return "", err
	}

	newGoalMap := reduceGoalMapToGoalIdList(budget.GoalMap)
	for _, val := range currentGoalMap {
		if (!utils.Contains(newGoalMap, val.Id)) {
			goal.RemoveBudgetIdFromGoal(ctx, val.Id, id)
		}
	}

	if len(budget.GoalMap) > 0 {
		for _, val := range budget.GoalMap {
			_, err := goal.UpdateBudgetIdsList(ctx, val.Id, id)
			if err != nil {
				return "", err
			}
		}
	}

	return db.UpdateBudgetRecordById(ctx, id, budget)
}

func DeleteBudgetById(ctx *gin.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteBudgetRecordById(ctx, id)
}

func reduceGoalMapToGoalIdList(goalMap []models.BudgetInputMap) []string {
	var result []string
	for _, val := range goalMap {
		result = append(result, val.Id)
	}
	return result 
}