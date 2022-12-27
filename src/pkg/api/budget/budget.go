package budget

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func getBudgets(ctx *gin.Context, budget *[]models.Budget) (error) {
	// TODO: input validation
	return db.GetAllBudgets(ctx, budget)
}

func getBudgetById(ctx *gin.Context, id string, budget *models.Budget) (error) {
	// TODO: input validation
	return db.GetBudgetRecordById(ctx, id, budget)
}

func createBudget(ctx *gin.Context, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.GetSavings()
	return db.InsertBudgetRecord(ctx, budget)
}

func updateBudgetById(ctx *gin.Context, id string, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.GetSavings()
	return db.UpdateBudgetRecordById(ctx, id, budget)
}

func deleteBudgetById(ctx *gin.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteBudgetRecordById(ctx, id)
}