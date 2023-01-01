package budget

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func GetBudgets(ctx context.Context, budget *[]models.Budget) (error) {
	// TODO: input validation
	return db.GetAllBudgets(ctx, budget)
}

func GetBudget(ctx context.Context, id string) (*models.Budget, error) {
	// TODO: input validation
	return db.GetBudgetRecordById(ctx, id)
}

func createBudget(ctx *gin.Context, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	return db.InsertBudgetRecord(ctx, budget)
}

func UpdateBudgetById(ctx *gin.Context, id string, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	return db.UpdateBudgetRecordById(ctx, id, budget)
}

func DeleteBudgetById(ctx *gin.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteBudgetRecordById(ctx, id)
}