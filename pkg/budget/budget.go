package budget

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/bill"
	"github.com/sanchitdeora/budget-tracker/pkg/goal"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/utils"
)

//go:generate mockgen -destination=./mocks/mock_budget.go -package=mock_budget github.com/sanchitdeora/budget-tracker/pkg/budget Service
type Service interface {
	GetBudgets(ctx context.Context) (*[]models.Budget, error)
	GetBudgetById(ctx context.Context, id string) (*models.Budget, error)
	GetGoalMap(ctx context.Context, id string) (*[]models.BudgetInputMap, error)
	CreateBudgetByUser(ctx context.Context, budget *models.Budget) (string, error)
	CreateBudget(ctx context.Context, budget *models.Budget, prevBudget *models.Budget) (string, error)
	UpdateBudgetById(ctx context.Context, id string, budget *models.Budget) (string, error)
	DeleteBudgetById(ctx context.Context, id string) (string, error)
}

type Opts struct {
	TransactionService 	transaction.Service
	BillService 		bill.Service
	GoalService			goal.Service
	DB					db.Database
}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}

func (s *serviceImpl) GetBudgets(ctx context.Context) (*[]models.Budget, error) {
	// TODO: input validation
	budgets, err := s.DB.GetAllBudgetRecords(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for i, budget := range *budgets {
		transactions, err := s.TransactionService.GetTransactionsByDate(ctx, budget.CreationTime, budget.ExpirationTime)
		if err != nil {
			return nil, err
		}

		updateCurrentAmounts(true, &(*budgets)[i].IncomeMap, transactions)
		updateCurrentAmounts(false, &(*budgets)[i].ExpenseMap, transactions)
	}

	return budgets, nil
}

func (s *serviceImpl) GetBudgetById(ctx context.Context, id string) (*models.Budget, error) {
	// TODO: input validation
	return s.DB.GetBudgetRecordById(ctx, id)
}

func (s *serviceImpl) CreateBudgetByUser(ctx context.Context, budget *models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.SetSavings()
	budget.SetByUser()
	setBudgetTime(budget)

	for _, val := range budget.GoalMap {
		s.GoalService.UpdateBudgetIdsList(ctx, val.Id, budget.BudgetId)
	}

	fmt.Println("Creating a new budget in service: ", budget)
	return s.DB.InsertBudgetRecord(ctx, budget)
}

func (s *serviceImpl) CreateBudget(ctx context.Context, budget *models.Budget, prevBudget *models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.SetSavings()
	budget.AutoSet(prevBudget.SequenceStartId, prevBudget.SequenceNumber)
	setBudgetTime(budget)

	fmt.Println("Creating a new budget: ", budget)
	return s.DB.InsertBudgetRecord(ctx, budget)
}

func (s *serviceImpl) UpdateBudgetById(ctx context.Context, id string, budget *models.Budget) (string, error) {
	// TODO: input validation
	budget.SetCategory()
	budget.SetFrequency()
	budget.SetSavings()

	currentGoalMap, err := s.GetGoalMap(ctx, id)
	if err != nil {
		return "", err
	}

	newGoalList := reduceGoalMapToGoalIdList(budget.GoalMap)

	for _, val := range *currentGoalMap {
		if (!utils.Contains(newGoalList, val.Id)) {
			s.GoalService.RemoveBudgetIdFromGoal(ctx, val.Id, id)
		}
	}

	if len(budget.GoalMap) > 0 {
		fmt.Println("Current budget: ", budget)
		for _, val := range budget.GoalMap {
			_, err := s.GoalService.UpdateBudgetIdsList(ctx, val.Id, id)
			if err != nil {
				return "", err
			}
		}
	}

	fmt.Println("Updating a new budget: ", budget)
	return s.DB.UpdateBudgetRecordById(ctx, id, budget)
}

func (s *serviceImpl) DeleteBudgetById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	budget, err := s.GetBudgetById(ctx, id)
	if err != nil {
		return "", err
	}

	for _, val := range budget.GoalMap {
		s.GoalService.RemoveBudgetIdFromGoal(ctx, val.Id, id)
	}

	return s.DB.DeleteBudgetRecordById(ctx, id)
}

// util functions

func filterTransactionsByCategory(categoryId string, transactions []models.Transaction) []models.Transaction {
	var filteredTransactions []models.Transaction

	for _, transaction := range transactions {
		if transaction.Category == categoryId {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	return filteredTransactions

}

func filterTransactionsByType(txType bool, transactions []models.Transaction) []models.Transaction {
	var filteredTransactions []models.Transaction

	for _, transaction := range transactions {
		if transaction.Type == txType {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	return filteredTransactions

}

func totalAmountOfTransactions(transactions []models.Transaction) float32 {
	var sum float32
	for _, transaction := range transactions {
		sum += transaction.Amount
	}
	return sum

}

func removeFilteredTransactions(filteredTransactions []models.Transaction, transactions *[]models.Transaction) {

	for _, fTx := range filteredTransactions {
		for i, tx := range *transactions {
			if fTx.TransactionId == tx.TransactionId {
				(*transactions)[i] = (*transactions)[len(*transactions) - 1]
			}
		}
		(*transactions) = (*transactions)[:len(*transactions) - 1]
	}

}

func (s *serviceImpl) GetGoalMap(ctx context.Context, id string) (*[]models.BudgetInputMap, error) {
	budget, err := s.GetBudgetById(ctx, id)
	return &budget.GoalMap, err
}

func reduceGoalMapToGoalIdList(goalMap []models.BudgetInputMap) []string {
	var result []string
	for _, val := range goalMap {
		result = append(result, val.Id)
	}
	return result 
}

func setBudgetTime(budget *models.Budget) (error) {
	if budget.CreationTime.IsZero() {
		budget.CreationTime = time.Now().Local()
	}

	expTime, err := utils.CalculateEndDateWithFrequency(budget.CreationTime, budget.Frequency)
	if err != nil {
		return err
	}
	budget.ExpirationTime = expTime
	return nil
}


func updateCurrentAmounts(transactionType bool, budgetMaps *[]models.BudgetInputMap, transactions *[]models.Transaction) {

	typeTransactions := filterTransactionsByType(transactionType, *transactions)

	uncategorizedMapIndex := -1
	var uncategorized models.BudgetInputMap

	uncategorized.Id = models.UNCATEGORIZED_CATEGORY
	uncategorized.Name = "Others"
	uncategorized.Amount = 0
	uncategorized.CurrentAmount = 0

	for i, budgetMap := range *budgetMaps {
		if budgetMap.Id != models.UNCATEGORIZED_CATEGORY {
			filteredTransactions := filterTransactionsByCategory(budgetMap.Id, typeTransactions)
			
			(*budgetMaps)[i].CurrentAmount = totalAmountOfTransactions(filteredTransactions)

			removeFilteredTransactions(filteredTransactions, &typeTransactions)
		} else {
			uncategorized.Amount = budgetMap.Amount
			uncategorizedMapIndex = i
		}
	}

	if len(typeTransactions) > 0 {
		if uncategorizedMapIndex == -1 {
			uncategorized.CurrentAmount = totalAmountOfTransactions(typeTransactions)
			*budgetMaps = append(*budgetMaps, uncategorized)
		} else {
			(*budgetMaps)[uncategorizedMapIndex].CurrentAmount = totalAmountOfTransactions(typeTransactions)
		}
	}
}