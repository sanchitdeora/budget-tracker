package goal

import (
	"context"
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/utils"
)

type Service interface {
	GetGoals(ctx context.Context, goal *[]models.Goal) (error)
	GetGoalById(ctx context.Context, id string) (*models.Goal, error)
	CreateGoalById(ctx context.Context, goal models.Goal) (string, error)
	UpdateGoalById(ctx context.Context, id string, goal models.Goal) (string, error)
	UpdateBudgetIdsList(ctx context.Context, goalId string, budgetId string) (string, error)
	RemoveBudgetIdFromGoal(ctx context.Context, goalId string, budgetId string) (string, error)
	DeleteGoalById(ctx context.Context, id string) (string, error)
}

type Opts struct {}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}

func (s *serviceImpl) GetGoals(ctx context.Context, goal *[]models.Goal) (error) {
	// TODO: input validation
	return db.GetAllGoals(ctx, goal)
}

func (s *serviceImpl) GetGoalById(ctx context.Context, id string) (*models.Goal, error) {
	// TODO: input validation
	return db.GetGoalRecordById(ctx, id)
}

func (s *serviceImpl) CreateGoalById(ctx context.Context, goal models.Goal) (string, error) {
	// TODO: input validation
	return db.InsertGoalRecord(ctx, goal)
}

func (s *serviceImpl) UpdateGoalById(ctx context.Context, id string, goal models.Goal) (string, error) {
	// TODO: input validation
	return db.UpdateGoalRecordById(ctx, id, goal)
}

func (s *serviceImpl) UpdateBudgetIdsList(ctx context.Context, goalId string, budgetId string) (string, error) {
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

func (s *serviceImpl) RemoveBudgetIdFromGoal(ctx context.Context, goalId string, budgetId string) (string, error) {
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

func (s *serviceImpl) DeleteGoalById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	
	return db.DeleteGoalRecordById(ctx, id)
}