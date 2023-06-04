package goal

import (
	"context"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/utils"
)

//go:generate mockgen -destination=../mocks/mock_goal.go -package=mocks github.com/sanchitdeora/budget-tracker/pkg/goal Service
type Service interface {
	GetGoals(ctx context.Context, goal *[]models.Goal) (error)
	GetGoalById(ctx context.Context, id string) (*models.Goal, error)
	CreateGoalById(ctx context.Context, goal models.Goal) (string, error)
	UpdateGoalById(ctx context.Context, id string, goal models.Goal) (string, error)
	UpdateBudgetIdsList(ctx context.Context, goalId string, budgetId string) (string, error)
	RemoveBudgetIdFromGoal(ctx context.Context, goalId string, budgetId string) (string, error)
	DeleteGoalById(ctx context.Context, id string) (string, error)
}

type Opts struct {
	DB db.Database
}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}

func (s *serviceImpl) GetGoals(ctx context.Context, goal *[]models.Goal) (error) {
	// TODO: input validation
	return s.DB.GetAllGoalRecords(ctx, goal)
}

func (s *serviceImpl) GetGoalById(ctx context.Context, id string) (*models.Goal, error) {
	// TODO: input validation
	return s.DB.GetGoalRecordById(ctx, id)
}

func (s *serviceImpl) CreateGoalById(ctx context.Context, goal models.Goal) (string, error) {
	// TODO: input validation
	goal.GoalId = db.GOAL_PREFIX + uuid.NewString()
	for _, val := range goal.BudgetIdList {
		s.updateGoalMapInBudget(ctx, val, goal)
	}

	return s.DB.InsertGoalRecord(ctx, goal)
}

func (s *serviceImpl) UpdateGoalById(ctx context.Context, id string, goal models.Goal) (string, error) {
	// TODO: input validation
	currGoal, err := s.GetGoalById(ctx, id)
	if err != nil {
		return "", nil
	}

	for _, val := range currGoal.BudgetIdList {
		if (!utils.Contains(goal.BudgetIdList, val)) {
			s.removeGoalMapInBudget(ctx, val, id)
		}
	}

	for _, val := range goal.BudgetIdList {
		if (!utils.Contains(currGoal.BudgetIdList, val)) {
			s.updateGoalMapInBudget(ctx, val, goal)
		}
	}

	return s.DB.UpdateGoalRecordById(ctx, id, goal)
}

func (s *serviceImpl) UpdateBudgetIdsList(ctx context.Context, goalId string, budgetId string) (string, error) {
	// TODO: input validation
	goal, err := s.DB.GetGoalRecordById(ctx, goalId)
	if err != nil {
		return "", nil
	}

	if (!utils.Contains(goal.BudgetIdList, budgetId)) {
		goal.BudgetIdList = append(goal.BudgetIdList, budgetId)	
	}
	return s.DB.UpdateGoalRecordById(ctx, goalId, *goal)
}

func (s *serviceImpl) RemoveBudgetIdFromGoal(ctx context.Context, goalId string, budgetId string) (string, error) {
	// TODO: input validation
	goal, err := s.DB.GetGoalRecordById(ctx, goalId)
	if err != nil {
		return "", nil
	}
	if (utils.Contains(goal.BudgetIdList, budgetId)) {
		index := utils.SearchIndex(goal.BudgetIdList, budgetId)
		goal.BudgetIdList = utils.Remove(goal.BudgetIdList, index)	
	}
	return s.DB.UpdateGoalRecordById(ctx, goalId, *goal)
}

func (s *serviceImpl) DeleteGoalById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	goal, err := s.GetGoalById(ctx, id)
	if err != nil {
		return "", err
	}

	for _, val := range goal.BudgetIdList {
		s.removeGoalMapInBudget(ctx, val, goal.GoalId)
	}
	
	return s.DB.DeleteGoalRecordById(ctx, id)
}

// using db directly to avoid cyclic imports. 
// TODO: Find better implementation to avoid this.

func (s *serviceImpl) removeGoalMapInBudget(ctx context.Context, budgetId string, goalId string) error {
	var goalList []string
	budget, _ := s.DB.GetBudgetRecordById(ctx, budgetId)
	
	for _, val := range budget.GoalMap {
		goalList = append(goalList, val.Id)
	}
	
	index := utils.SearchIndex(goalList, goalId)

	budget.GoalMap = append(budget.GoalMap[:index], budget.GoalMap[index+1:]...)

	s.DB.UpdateBudgetRecordById(ctx, budgetId, *budget)
	return nil
}

func (s *serviceImpl) updateGoalMapInBudget(ctx context.Context, budgetId string, goal models.Goal) error {
	budget, _ := s.DB.GetBudgetRecordById(ctx, budgetId)
	budget.GoalMap = append(budget.GoalMap, models.BudgetInputMap{Id: goal.GoalId, Name: goal.GoalName, Amount: 0})
	
	s.DB.UpdateBudgetRecordById(ctx, budgetId, *budget)
	return nil
}