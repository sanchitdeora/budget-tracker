package goal

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/sanchitdeora/budget-tracker/utils"
)

//go:generate mockgen -destination=./mocks/mock_goal.go -package=mock_goal github.com/sanchitdeora/budget-tracker/pkg/goal Service
type Service interface {
	GetGoals(ctx context.Context) (*[]models.Goal, error)
	GetGoalById(ctx context.Context, id string) (*models.Goal, error)
	// GetCurrentAmountInGoals(ctx context.Context, goal *models.Goal) (float32, error)
	CreateGoal(ctx context.Context, goal *models.Goal) (string, error)
	UpdateGoalById(ctx context.Context, id string, goal *models.Goal) (string, error)
	// UpdateGoalAmount(ctx context.Context, goalId string, currAmount float32, budgetId string) (string, error)
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

func (s *serviceImpl) GetGoals(ctx context.Context) (*[]models.Goal, error) {
	goals, err := s.DB.GetAllGoalRecords(ctx)
	if err != nil {
		return nil, err
	}

	for idx := range *goals {
		if err = s.updateCurrentAmountInGoals(ctx, &(*goals)[idx]); err != nil {
			return nil, err
		}
	}
	return goals, nil
}

func (s *serviceImpl) GetGoalById(ctx context.Context, id string) (*models.Goal, error) {
	if id == "" {
		log.Println("missing Goal Id", ctx)
		return nil, exceptions.ErrValidationError
	}

	goal, err := s.DB.GetGoalRecordById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = s.updateCurrentAmountInGoals(ctx, goal); err != nil {
		return nil, err
	}

	return goal, nil
}

func (s *serviceImpl) CreateGoal(ctx context.Context, goal *models.Goal) (string, error) {
	if !goal.IsValid() {
		return "", exceptions.ErrValidationError
	}

	goal.GoalId = db.GOAL_PREFIX + uuid.NewString()
	for _, val := range goal.BudgetIdList {
		if err := s.updateGoalMapInBudget(ctx, val, goal); err != nil {
			return "", err
		}
	}

	return s.DB.InsertGoalRecord(ctx, goal)
}

func (s *serviceImpl) UpdateGoalById(ctx context.Context, id string, goal *models.Goal) (string, error) {
	if id == "" {
		log.Println("missing Goal Id", ctx)
		return "", exceptions.ErrValidationError
	}

	if !goal.IsValid() {
		return "", exceptions.ErrValidationError
	}

	currGoal, err := s.GetGoalById(ctx, id)
	if err != nil {
		return "", err
	}

	for _, val := range currGoal.BudgetIdList {
		if (!utils.Contains(goal.BudgetIdList, val)) {
			if err = s.removeGoalMapInBudget(ctx, val, id); err != nil {
				return "", err
			}
		}
	}

	for _, val := range goal.BudgetIdList {
		if (!utils.Contains(currGoal.BudgetIdList, val)) {
			if err := s.updateGoalMapInBudget(ctx, val, goal); err != nil {
				return "", err
			}
		}
	}

	return s.DB.UpdateGoalRecordById(ctx, id, goal)
}

func (s *serviceImpl) UpdateBudgetIdsList(ctx context.Context, goalId string, budgetId string) (string, error) {
	if goalId == "" {
		log.Println("missing Goal Id", ctx)
		return "", exceptions.ErrValidationError
	}
	if budgetId == "" {
		log.Println("missing Budget Id", ctx)
		return "", exceptions.ErrValidationError
	}

	goal, err := s.GetGoalById(ctx, goalId)
	if err != nil {
		return "", err
	}

	if (!utils.Contains(goal.BudgetIdList, budgetId)) {
		goal.BudgetIdList = append(goal.BudgetIdList, budgetId)	
		return s.DB.UpdateGoalRecordById(ctx, goalId, goal)
	}
	return goalId, nil
}

func (s *serviceImpl) RemoveBudgetIdFromGoal(ctx context.Context, goalId string, budgetId string) (string, error) {
	if goalId == "" {
		log.Println("missing Goal Id", ctx)
		return "", exceptions.ErrValidationError
	}
	if budgetId == "" {
		log.Println("missing Budget Id", ctx)
		return "", exceptions.ErrValidationError
	}

	goal, err := s.GetGoalById(ctx, goalId)
	if err != nil {
		return "", err
	}
	if utils.Contains(goal.BudgetIdList, budgetId) {
		index := utils.SearchIndex(goal.BudgetIdList, budgetId)
		goal.BudgetIdList = utils.Remove(goal.BudgetIdList, index)	
		return s.DB.UpdateGoalRecordById(ctx, goalId, goal)
	}
	return goalId, nil
}

func (s *serviceImpl) DeleteGoalById(ctx context.Context, id string) (string, error) {
	if id == "" {
		return "", exceptions.ErrValidationError
	}

	goal, err := s.GetGoalById(ctx, id)
	if err != nil {
		return "", err
	}

	for _, val := range goal.BudgetIdList {
		err = s.removeGoalMapInBudget(ctx, val, goal.GoalId)
		if err != nil {
			return "", err
		}
	}

	return s.DB.DeleteGoalRecordById(ctx, id)
}

// using db directly to avoid cyclic imports. 
// TODO: Find better implementation to avoid this.
func (s *serviceImpl) removeGoalMapInBudget(ctx context.Context, budgetId string, goalId string) error {
	var goalList []string
	budget, err := s.DB.GetBudgetRecordById(ctx, budgetId)
	if err != nil {
		return err
	}

	for _, val := range budget.GoalMap {
		goalList = append(goalList, val.Id)
	}

	index := utils.SearchIndex(goalList, goalId)

	budget.GoalMap = append(budget.GoalMap[:index], budget.GoalMap[index+1:]...)

	_, err = s.DB.UpdateBudgetRecordById(ctx, budgetId, budget)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceImpl) updateGoalMapInBudget(ctx context.Context, budgetId string, goal *models.Goal) error {
	budget, err := s.DB.GetBudgetRecordById(ctx, budgetId)
	if err != nil {
		return err
	}

	budget.GoalMap = append(budget.GoalMap, models.BudgetInputMap{Id: goal.GoalId, Name: goal.GoalName, Amount: 0})

	_, err = s.DB.UpdateBudgetRecordById(ctx, budgetId, budget)
	if err != nil {
		log.Println("error while updating goal map in budget, goalId:", goal.GoalId, "budgetId:", budgetId, "error:", err)
		return err
	}
	return nil
}

func (s *serviceImpl) updateCurrentAmountInGoals(ctx context.Context, goal *models.Goal) error {
	amount, err := s.getCurrentAmountInGoals(ctx, goal)
	if err != nil {
		return err
	}

	(*goal).CurrentAmount += amount
	return nil
}

func (s *serviceImpl) getCurrentAmountInGoals(ctx context.Context, goal *models.Goal) (float32, error) {
	var currAmount float32 = 0
	fmt.Println("Current goal in GetCurrentAmountInGoals:", goal)

	for _, budgetId := range goal.BudgetIdList {
		budget, err := s.DB.GetBudgetRecordById(ctx, budgetId)
		if err != nil {
			log.Println("error while fetching budgets", err, ctx)
			return currAmount, err
		}
		if budget == nil {
			log.Println("no Budget record found for budgetId:", budgetId, ctx)
			return 0, exceptions.ErrNoBudgetsFound
		}

		// fmt.Println("Current budget in GetCurrentAmountInGoals:", budget)
		for _, bGoal := range (*budget).GoalMap {
			if bGoal.Id == goal.GoalId {
				currAmount += bGoal.CurrentAmount
			}
		}
	}

	return currAmount, nil
}

// Need to figure out where its used
// func (s *serviceImpl) UpdateGoalAmount(ctx context.Context, goalId string, currAmount float32, budgetId string) (string, error) {
// 	if goalId == "" {
// 		log.Println("missing Goal Id", ctx)
// 		return "", exceptions.ErrValidationError
// 	}
// 	if budgetId == "" {
// 		log.Println("missing Budget Id", ctx)
// 		return "", exceptions.ErrValidationError
// 	}

// 	goal, err := s.DB.GetGoalRecordById(ctx, goalId)
// 	if err != nil {
// 		return "", err
// 	}

// 	if currAmount > 0 {
// 		if currAmount >= goal.TargetAmount {
// 			goal.CurrentAmount = goal.TargetAmount
// 		} else {
// 			goal.CurrentAmount += currAmount
// 		}
// 	}

// 	if (!utils.Contains(goal.BudgetIdList, budgetId)) {
// 		goal.BudgetIdList = append(goal.BudgetIdList, budgetId)	
// 	}
// 	return s.DB.UpdateGoalRecordById(ctx, goalId, goal)
// }