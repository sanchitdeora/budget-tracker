package budget

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/bill"
	"github.com/sanchitdeora/budget-tracker/pkg/goal"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/sanchitdeora/budget-tracker/utils"
)

//go:generate mockgen -destination=./mocks/mock_budget.go -package=mock_budget github.com/sanchitdeora/budget-tracker/pkg/budget Service
type Service interface {
	GetBudgets(ctx context.Context) (*[]models.Budget, error)
	GetBudgetById(ctx context.Context, id string) (*models.Budget, error)
	CreateBudgetByUser(ctx context.Context, budget *models.Budget) (string, error)
	CreateRecurringBudget(ctx context.Context, budget *models.Budget, prevBudget *models.Budget) (string, error)
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
	budgets, err := s.DB.GetAllBudgetRecords(ctx)
	if err != nil {
		log.Println("error getting all budget records", ctx, err)
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
	if id == "" {
		log.Println("missing Budget Id", ctx)
		return nil, exceptions.ErrValidationError
	}
	return s.DB.GetBudgetRecordById(ctx, id)
}

func (s *serviceImpl) CreateBudgetByUser(ctx context.Context, budget *models.Budget) (string, error) {
	if err := validateBudget(ctx, budget); err != nil {
		return "", err
	}
	budget.SetByUser()
	setBudgetTime(budget)

	for _, val := range budget.GoalMap {
		_, err := s.GoalService.UpdateBudgetIdsList(ctx, val.Id, budget.BudgetId)
		if err != nil {
			log.Println("error while updating budgets ids list for the goal:", val.Id, "budgetId:", budget.BudgetId, "error:", err, ctx)
			return "", err
		}
	}

	log.Println("creating a new budget by user request: ", budget, ctx)
	return s.DB.InsertBudgetRecord(ctx, budget)
}

// Plan this more
func (s *serviceImpl) CreateRecurringBudget(ctx context.Context, budget *models.Budget, prevBudget *models.Budget) (string, error) {
	if err := validateBudget(ctx, budget); err != nil {
		return "", err
	}
	if prevBudget.BudgetId == "" {
		log.Println("missing Budget Id while creating a recurring budget", ctx)
		return "", exceptions.ErrValidationError
	}
	if !prevBudget.IsValid() {
		log.Println("prev budget:",prevBudget.BudgetId, "is invalid", ctx)
		return "", exceptions.ErrValidationError
	}

	budget.AutoSet(prevBudget.SequenceStartId, prevBudget.SequenceNumber)
	setBudgetTime(budget)

	log.Println("creating a new budget: ", budget, ctx)

	// update prev budget to be marked as closed
	// prevBudget.IsClosed = true
	// s.DB.UpdateBudgetRecordById(ctx, prevBudget.BudgetId, prevBudget)

	return s.DB.InsertBudgetRecord(ctx, budget)
}

func (s *serviceImpl) UpdateBudgetById(ctx context.Context, id string, budget *models.Budget) (string, error) {
	if id == "" {
		log.Println("missing Budget Id", ctx)
		return "", exceptions.ErrValidationError
	}
	if err := validateBudget(ctx, budget); err != nil {
		return "", err
	}

	currentBudget, err := s.GetBudgetById(ctx, id)
	if err != nil {
		return "", err
	}

	newGoalList := reduceGoalMapToGoalIdList(budget.GoalMap)

	for _, val := range currentBudget.GoalMap {
		if (!utils.Contains(newGoalList, val.Id)) {
			_, err := s.GoalService.RemoveBudgetIdFromGoal(ctx, val.Id, id)
			if err != nil {
				log.Println("error while removing budgets ids list for the goal:", val.Id, "budgetId:", budget.BudgetId, "error:", err, ctx)
				return "", err
			}
		}
	}

	if len(budget.GoalMap) > 0 {
		for _, val := range budget.GoalMap {
			_, err := s.GoalService.UpdateBudgetIdsList(ctx, val.Id, id)
			if err != nil {
				log.Println("error while updating budgets ids list for the goal:", val.Id, "budgetId:", budget.BudgetId, "error:", err, ctx)
				return "", err
			}
		}
	}

	log.Println("Updating a new budget: ", budget)
	return s.DB.UpdateBudgetRecordById(ctx, id, budget)
}

func (s *serviceImpl) DeleteBudgetById(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("missing Budget Id while deleting budget", ctx)
		return "", exceptions.ErrValidationError
	}
	budget, err := s.GetBudgetById(ctx, id)
	if err != nil {
		return "", err
	}

	for _, val := range budget.GoalMap {
		_, err = s.GoalService.RemoveBudgetIdFromGoal(ctx, val.Id, id)
		if err != nil {
			log.Println("error while removing budget id from goal:", val.Id, "budgetId:", id, "error:", err, ctx)
			return "", err
		}
	}

	return s.DB.DeleteBudgetRecordById(ctx, id)
}

// util functions

func validateBudget(ctx context.Context, budget *models.Budget) error{
	if !budget.IsValid() {
		return exceptions.ErrValidationError
	}
	
	budget.SetCategory()
	budget.SetFrequency()
	budget.SetSavings()

	return nil
}

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

func reduceGoalMapToGoalIdList(goalMap []models.BudgetInputMap) []string {
	var result []string
	for _, val := range goalMap {
		result = append(result, val.Id)
	}
	return result 
}

func setBudgetTime(budget *models.Budget) {
	if budget.CreationTime.IsZero() {
		budget.CreationTime = time.Now().Local()
	}

	expTime := utils.CalculateEndDateWithFrequency(budget.CreationTime, budget.Frequency)
	budget.ExpirationTime = expTime
}

func updateCurrentAmounts(transactionType bool, budgetMaps *[]models.BudgetInputMap, transactions *[]models.Transaction) {

	filteredTypeTransactions := filterTransactionsByType(transactionType, *transactions)

	uncategorizedMapIndex := -1
	var uncategorized models.BudgetInputMap

	uncategorized.Id = models.UNCATEGORIZED_CATEGORY
	uncategorized.Name = "Others"
	uncategorized.Amount = 0
	uncategorized.CurrentAmount = 0

	for i, budgetMap := range *budgetMaps {
		if budgetMap.Id != models.UNCATEGORIZED_CATEGORY {
			filteredCategoryTransactions := filterTransactionsByCategory(budgetMap.Id, filteredTypeTransactions)
			
			(*budgetMaps)[i].CurrentAmount = totalAmountOfTransactions(filteredCategoryTransactions)

			removeFilteredTransactions(filteredCategoryTransactions, &filteredTypeTransactions)
		} else {
			uncategorized.Amount = budgetMap.Amount
			uncategorizedMapIndex = i
		}
	}

	if len(filteredTypeTransactions) > 0 {
		if uncategorizedMapIndex == -1 {
			uncategorized.CurrentAmount = totalAmountOfTransactions(filteredTypeTransactions)
			*budgetMaps = append(*budgetMaps, uncategorized)
		} else {
			(*budgetMaps)[uncategorizedMapIndex].CurrentAmount = totalAmountOfTransactions(filteredTypeTransactions)
		}
	}
}