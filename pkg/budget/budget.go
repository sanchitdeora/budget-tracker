package budget

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/bill"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/sanchitdeora/budget-tracker/pkg/goal"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/utils"
)

//go:generate mockgen -destination=./mocks/mock_budget.go -package=mock_budget github.com/sanchitdeora/budget-tracker/pkg/budget Service
type Service interface {
	GetBudgets(ctx context.Context) (*[]models.Budget, error)
	GetBudgetById(ctx context.Context, id string) (*models.Budget, error)
	CreateBudgetByUser(ctx context.Context, budget *models.Budget) (string, error)
	UpdateBudgetById(ctx context.Context, id string, budget *models.Budget) (string, error)
	UpdateBudgetIsClosed(ctx context.Context, id string, isClosed bool) (string, error)
	DeleteBudgetById(ctx context.Context, id string) (string, error)
	
	// Maintain recurring budgets
	BudgetMaintainer(ctx context.Context)
	CreateRecurringBudget(ctx context.Context, prevBudget *models.Budget) (string, error)
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

	for i := range *budgets {
		err = s.updateBudgetCurrentAmount(ctx, &(*budgets)[i])
		if err != nil {
			return nil, err
		}
	}

	return budgets, nil
}

func (s *serviceImpl) GetBudgetById(ctx context.Context, id string) (*models.Budget, error) {
	if id == "" {
		log.Println("missing Budget Id", ctx)
		return nil, exceptions.ErrValidationError
	}

	budget, err := s.DB.GetBudgetRecordById(ctx, id)
	if err != nil {
		log.Println("error getting budget record for id:", id, ctx, err)
		return nil, err
	}

	err = s.updateBudgetCurrentAmount(ctx, budget)
	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (s *serviceImpl) CreateBudgetByUser(ctx context.Context, budget *models.Budget) (string, error) {
	if !budget.IsValid() {
		log.Println("budget is invalid:", budget)
		return "", exceptions.ErrValidationError
	}
	budget.SetCategory()
	budget.SetFrequency()

	setBudgetTime(budget)
	
	log.Println("creating a new budget by user request:", budget, ctx)
	
	budgetId, err := s.DB.InsertBudgetRecord(ctx, budget)
	if err != nil {
		log.Println("error while inserting budget:", budget, "error:", err, ctx)
		return "", err
	}

	for _, val := range budget.GoalMap {
		_, err := s.GoalService.UpdateBudgetIdsList(ctx, val.Id, budgetId)
		if err != nil {
			log.Println("error while updating budgets ids list for the goal:", val.Id, "budgetId:", budgetId, "error:", err, ctx)
			return "", err
		}
	}

	return budgetId, nil
}

// Plan this more
func (s *serviceImpl) CreateRecurringBudget(ctx context.Context, prevBudget *models.Budget) (string, error) {
	if prevBudget == nil || prevBudget.BudgetId == "" {
		log.Println("missing Budget Id while creating a recurring budget", ctx)
		return "", exceptions.ErrValidationError
	}
	if !prevBudget.IsValid() {
		log.Println("budget:", prevBudget.BudgetId, "is invalid", ctx)
		return "", exceptions.ErrValidationError
	}

	var budget models.Budget
	budget.CreateFromPreviousBudget(prevBudget)
	setBudgetTime(&budget)

	log.Println("creating a new budget:", budget, ctx)

	return s.DB.InsertBudgetRecord(ctx, &budget)
}

func (s *serviceImpl) UpdateBudgetById(ctx context.Context, id string, budget *models.Budget) (string, error) {
	if id == "" {
		log.Println("missing Budget Id", ctx)
		return "", exceptions.ErrValidationError
	}
	if !budget.IsValid() {
		return "", exceptions.ErrValidationError
	}
	setBudgetTime(budget)

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

	log.Println("updating a new budget:", budget, ctx)
	return s.DB.UpdateBudgetRecordById(ctx, id, budget)
}

func (s *serviceImpl) UpdateBudgetIsClosed(ctx context.Context, id string, isClosed bool) (string, error) {
	if id == "" {
		log.Println("missing Budget Id", ctx)
		return "", exceptions.ErrValidationError
	}

	budget, err := s.GetBudgetById(ctx, id)
	if err != nil {
		return "", err
	}

	budget.IsClosed = isClosed
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

// Maintainer functions
func (s *serviceImpl) BudgetMaintainer(ctx context.Context) {
	log.Println("start budget maintainer", ctx)

	successfulCounter := 0
	budgets, err := s.GetBudgets(ctx)
	if err != nil {
		log.Println("error while fetching budget records", "error:", err, ctx)
		return
	}
	
	for _, budget := range *budgets{
		if budget.Frequency == models.ONCE_FREQUENCY || 
			budget.NextSequenceId != "" || 
			budget.ExpirationTime.IsZero() || 
			time.Since(budget.ExpirationTime) < 0 {

			continue
		}

		if time.Since(budget.ExpirationTime) >= 0 {
			newBudgetId, err := s.CreateRecurringBudget(ctx, &budget)
			if err != nil {
				log.Println("error while creating new budget id:", newBudgetId, "error:", err, ctx)
				continue
			}
			
			budget.NextSequenceId = newBudgetId
			budget.IsClosed = true
			_, err = s.DB.UpdateBudgetRecordById(ctx, budget.BudgetId, &budget)
			if err != nil {
				log.Println("error while updating budget id:", budget.BudgetId, "error:", err, ctx)
				continue
			}
			successfulCounter++
		}
	}

	log.Println("complete budget maintainer with successfully updating", successfulCounter, "budgets", ctx)
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
	budget.ExpirationTime = utils.CalculateEndDateWithFrequency(budget.CreationTime, budget.Frequency)
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

func (s *serviceImpl) updateBudgetCurrentAmount(ctx context.Context, budget *models.Budget) error {
	transactions, err := s.TransactionService.GetTransactionsByDate(ctx, budget.CreationTime, budget.ExpirationTime)
	if err != nil {
		log.Println("error while fetching transactionsByDate for budgetId:", budget.BudgetId, err, ctx)
		return  err
	}

	if transactions != nil && len(*transactions) > 0 {
		updateCurrentAmounts(true, &budget.IncomeMap, transactions)
		updateCurrentAmounts(false, &budget.ExpenseMap, transactions)
	} else {
		resetBudgetMapCurrentAmount(&budget.IncomeMap)
		resetBudgetMapCurrentAmount(&budget.ExpenseMap)
	}

	return nil
}

func resetBudgetMapCurrentAmount(budgetMaps *[]models.BudgetInputMap) {
	if budgetMaps == nil || len(*budgetMaps) <= 0 {
		return
	}

	for i := range *budgetMaps {
		(*budgetMaps)[i].CurrentAmount = 0
	}
}