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
	GetGoals(ctx context.Context, goal *[]models.Goal) (error)
	GetGoalById(ctx context.Context, id string) (*models.Goal, error)
	GetCurrentAmountInGoals(ctx context.Context, goal *models.Goal) (float32, error)
	CreateGoalById(ctx context.Context, goal models.Goal) (string, error)
	UpdateGoalById(ctx context.Context, id string, goal models.Goal) (string, error)
	UpdateGoalAmount(ctx context.Context, goalId string, currAmount float32, budgetId string) (string, error)
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

func (s *serviceImpl) GetGoals(ctx context.Context, goals *[]models.Goal) (error) {
	// TODO: input validation
	err := s.DB.GetAllGoalRecords(ctx, goals)
	if err != nil {
		return err
	}

	for idx := range *goals {
		amount, err := s.GetCurrentAmountInGoals(ctx, &(*goals)[idx])
		if err != nil {
			return err
		}

		(*goals)[idx].CurrentAmount += amount
	}
	return nil
}

func (s *serviceImpl) GetGoalById(ctx context.Context, id string) (*models.Goal, error) {
	// TODO: input validation
	goal, err := s.DB.GetGoalRecordById(ctx, id)
	if err != nil {
		return nil, err
	}

	amount, err := s.GetCurrentAmountInGoals(ctx, goal)
	if err != nil {
		return nil, err
	}

	(*goal).CurrentAmount += amount

	return goal, nil
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
	log.Println("goal here1:", goal)
	if err != nil {
		return "", err
	}

	log.Println("goal here2:", goal)
	if (!utils.Contains(goal.BudgetIdList, budgetId)) {
		goal.BudgetIdList = append(goal.BudgetIdList, budgetId)	
	}
	return s.DB.UpdateGoalRecordById(ctx, goalId, *goal)
}

func (s *serviceImpl) UpdateGoalAmount(ctx context.Context, goalId string, currAmount float32, budgetId string) (string, error) {
	// TODO: input validation
	goal, err := s.DB.GetGoalRecordById(ctx, goalId)
	log.Println("goal here1:", goal)
	if err != nil {
		return "", err
	}

	if currAmount > 0 {
		if currAmount >= goal.TargetAmount {
			goal.CurrentAmount = goal.TargetAmount
		} else {
			goal.CurrentAmount += currAmount
		}
	}

	log.Println("goal here2:", goal)
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

func (s *serviceImpl) GetCurrentAmountInGoals(ctx context.Context, goal *models.Goal) (float32, error) {
	var currAmount float32 = 0
	fmt.Println("Current goal in GetCurrentAmountInGoals: ", goal)

	for _, budgetId := range goal.BudgetIdList {
		budget, err := s.DB.GetBudgetRecordById(ctx, budgetId)
		if budget == nil {
			log.Println("No Budget record found")
			return 0, exceptions.ErrNoBudgetsFound
		}
		if err != nil {
			log.Fatal("Error while fetching budgets", err)
			return currAmount, err
		}

		fmt.Println("Current budget in GetCurrentAmountInGoals: ", budget)
		for _, bGoal := range (*budget).GoalMap {
			if bGoal.Id == goal.GoalId {
				currAmount += bGoal.CurrentAmount
			}
		}
	}

	return currAmount, nil
}