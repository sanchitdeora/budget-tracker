package goal

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
	"github.com/sanchitdeora/budget-tracker/src/utils"
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

func UpdateBudgetIdsList(ctx *gin.Context, goalId string, budgetId string) (string, error) {
	// TODO: input validation
	goal, err := db.GetGoalRecordById(ctx, goalId)
	log.Print("\ngoal here: ", *goal)
	if err != nil {
		log.Fatal(err)
	}

	if (!utils.Contains(goal.BudgetIdList, budgetId)) {
		goal.BudgetIdList = append(goal.BudgetIdList, budgetId)	
	}
	return db.UpdateGoalRecordById(ctx, goalId, *goal)
}

func RemoveBudgetIdFromGoal(ctx *gin.Context, goalId string, budgetId string) (string, error) {
	// TODO: input validation
	goal, err := db.GetGoalRecordById(ctx, goalId)
	if err != nil {
		log.Fatal(err)
	}
	if (utils.Contains(goal.BudgetIdList, budgetId)) {
		index := utils.SearchIndex(goal.BudgetIdList, budgetId)
		goal.BudgetIdList = utils.Remove(goal.BudgetIdList, index)	
	}
	return db.UpdateGoalRecordById(ctx, goalId, *goal)
}

func DeleteGoalById(ctx *gin.Context, id string) (string, error) {
	// TODO: input validation
	
	return db.DeleteGoalRecordById(ctx, id)
}