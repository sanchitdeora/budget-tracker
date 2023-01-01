package goals

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func GetGoals(ctx *gin.Context, goal *[]models.Goal) (error) {
	// TODO: input validation
	return db.GetAllGoals(ctx, goal)
}

func GetGoal(ctx *gin.Context, id string) (*models.Goal, error) {
	// TODO: input validation
	return db.GetGoalRecordById(ctx, id)
}

func createGoal(ctx *gin.Context, goal models.Goal) (string, error) {
	// TODO: input validation
	return db.InsertGoalRecord(ctx, goal)
}

func UpdateGoalById(ctx *gin.Context, id string, goal models.Goal) (string, error) {
	// TODO: input validation
	return db.UpdateGoalRecordById(ctx, id, goal)
}

func DeleteGoalById(ctx *gin.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteGoalRecordById(ctx, id)
}