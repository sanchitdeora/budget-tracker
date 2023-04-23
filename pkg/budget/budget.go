package budget

import (
	"context"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/bill"
	"github.com/sanchitdeora/budget-tracker/pkg/goal"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/utils"
)

type Service interface {
	GetBudgets(ctx context.Context, budget *[]models.Budget) (error)
	GetBudgetById(ctx context.Context, id string) (*models.Budget, error)
	GetGoalMap(ctx context.Context, id string) ([]models.BudgetInputMap, error)
	CreateBudgetByUser(ctx context.Context, budget models.Budget) (string, error)
	CreateBudget(ctx context.Context, budget models.Budget) (string, error)
	UpdateBudgetById(ctx context.Context, id string, budget models.Budget) (string, error)
	DeleteBudgetById(ctx context.Context, id string) (string, error)
}

type Opts struct {
	TransactionService 	transaction.Service
	BillService 		bill.Service
	GoalService			goal.Service
}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}

func (s *serviceImpl) GetBudgets(ctx context.Context, budget *[]models.Budget) (error) {
	// TODO: input validation
	return db.GetAllBudgets(ctx, budget)
}

func (s *serviceImpl) GetBudgetById(ctx context.Context, id string) (*models.Budget, error) {
	// TODO: input validation
	return db.GetBudgetRecordById(ctx, id)
}

func (s *serviceImpl) GetGoalMap(ctx context.Context, id string) ([]models.BudgetInputMap, error) {
	budget, err := s.GetBudgetById(ctx, id)
	return budget.GoalMap, err
}

func (s *serviceImpl) CreateBudgetByUser(ctx context.Context, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	budget.SetByUser()
	return db.InsertBudgetRecord(ctx, budget)
}

func (s *serviceImpl) CreateBudget(ctx context.Context, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	budget.AutoSet()
	return db.InsertBudgetRecord(ctx, budget)
}

func (s *serviceImpl) UpdateBudgetById(ctx context.Context, id string, budget models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.GetSavings()
	
	currentGoalMap, err := s.GetGoalMap(ctx, id)
	if err != nil {
		return "", err
	}

	newGoalMap := reduceGoalMapToGoalIdList(budget.GoalMap)
	for _, val := range currentGoalMap {
		if (!utils.Contains(newGoalMap, val.Id)) {
			s.GoalService.RemoveBudgetIdFromGoal(ctx, val.Id, id)
		}
	}

	if len(budget.GoalMap) > 0 {
		for _, val := range budget.GoalMap {
			_, err := s.GoalService.UpdateBudgetIdsList(ctx, val.Id, id)
			if err != nil {
				return "", err
			}
		}
	}

	return db.UpdateBudgetRecordById(ctx, id, budget)
}

func (s *serviceImpl) DeleteBudgetById(ctx context.Context, id string) (string, error) {
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