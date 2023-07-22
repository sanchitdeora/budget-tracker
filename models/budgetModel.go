package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Budget struct {
	BudgetId   		string           `json:"budget_id"`
	BudgetName		string           `json:"budget_name"`
	IncomeMap		[]BudgetInputMap `json:"income_map"`
	ExpenseMap 		[]BudgetInputMap `json:"expense_map"`
	GoalMap    		[]BudgetInputMap `json:"goal_map"`
	Frequency   	string           `json:"frequency"`
	Savings     	float32          `json:"savings"`
	CreationTime    time.Time 		 `json:"creation_time"`
	ExpirationTime  time.Time 		 `json:"expiration_time"`
	SequenceStartId string    		 `json:"sequence_start_id"`
	SequenceNumber  int       		 `json:"sequence_no"`
	IsClosed 		bool	    	 `json:"is_closed"`
}

type BudgetInputMap struct  {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	CurrentAmount float32 `json:"current_amount" bson:"current_amount"`
	Amount        float32 `json:"amount"`
}

var BudgetFrequencyMap = []string{
	ONCE_FREQUENCY,
	DAILY_FREQUENCY,
	WEEKLY_FREQUENCY,
	BI_WEEKLY_FREQUENCY,
	MONTHLY_FREQUENCY,
	BI_MONTHLY_FREQUENCY,
	QUARTERLY_FREQUENCY,
	HALF_YEARLY_FREQUENCY,
	YEARLY_FREQUENCY,
}

const BUDGET_PREFIX = "BG-"

func (budget *Budget) SetSavings() float32 {
	budget.Savings = calculateSavings(budget.IncomeMap, budget.ExpenseMap, budget.GoalMap)
	log.Print("set savings here:", budget.Savings)
	return budget.Savings
}

func (budget *Budget) SetFrequency() {
	for index, frequency := range BudgetFrequencyMap {
		if frequency == strings.ToLower(budget.Frequency) {
			budget.Frequency = BudgetFrequencyMap[index]
			return
		}
		// default frequency set as once
		if index == len(BudgetFrequencyMap) - 1 {
			budget.Frequency = ONCE_FREQUENCY
		}
	}
}

func (budget *Budget) SetCategory() {
	// replace unknown categories with uncategorized
	budget.IncomeMap = updateCategoryInBudgetMap(budget.IncomeMap)
	budget.ExpenseMap = updateCategoryInBudgetMap(budget.ExpenseMap)

	fmt.Println("Main Object expense map: ", budget.ExpenseMap)
}

func (budget *Budget) SetByUser() {
	budgetId := BUDGET_PREFIX + uuid.NewString()
	budget.BudgetId = budgetId

	budget.SequenceStartId = budgetId
	budget.SequenceNumber = 0

}
func (budget *Budget) AutoSet(sequenceStardId string, prevSequenceNo int) {
	budgetId := BUDGET_PREFIX + uuid.NewString()
	budget.BudgetId = budgetId
	
	budget.SequenceNumber = prevSequenceNo + 1
	budget.SequenceStartId = sequenceStardId

}

func updateCategoryInBudgetMap(budgetMap []BudgetInputMap) []BudgetInputMap{
	var uncategorized BudgetInputMap
	uncategorized.Id = UNCATEGORIZED_CATEGORY
	uncategorized.Name = "Uncategorized"
	uncategorized.Amount = 0

	var valToBeDeleted []string

	for _, val := range budgetMap {
		for index, allowedCategory := range TransactionCategoryMap {
			if allowedCategory == strings.ToLower(val.Id) {
				break
			}
			if index == len(TransactionCategoryMap) - 1 {
				fmt.Println("allowedCategory here: ", allowedCategory)

				uncategorized.Amount += val.Amount
				valToBeDeleted = append(valToBeDeleted, val.Id)
			}
		}
	}
	if len(valToBeDeleted) > 0 {
		fmt.Println("values to be deleted", valToBeDeleted)

		for _, id := range valToBeDeleted {
			i := slices.IndexFunc(budgetMap, func(b BudgetInputMap) bool { return b.Id == id })

			budgetMap[i] = budgetMap[len(budgetMap) - 1]
			budgetMap = append(budgetMap[:i], budgetMap[i+1:]...)

			fmt.Println(budgetMap)
		}

		fmt.Println("after deleting:", budgetMap)
	}

	if uncategorized.Amount > 0 {
		idx := slices.IndexFunc(budgetMap, func(b BudgetInputMap) bool { return b.Id == UNCATEGORIZED_CATEGORY })
		if idx > -1 {
			fmt.Println("index for uncategorized found: ", idx, "with value: ", budgetMap[idx])
			budgetMap[idx].Amount += uncategorized.Amount
		} else {
			budgetMap = append(budgetMap, uncategorized)
		}
	}
	fmt.Println("after uncategorized setting:", budgetMap)

	return budgetMap
}

func calculateSavings(incomeMap []BudgetInputMap, spendingLimitMap []BudgetInputMap, goalAmountMap []BudgetInputMap) float32 {
	var totalSavings float32
	for _, val := range incomeMap {
		totalSavings += val.Amount
	}

	if len(spendingLimitMap) != 0 {
		for _, val := range spendingLimitMap {
			totalSavings -= val.Amount
		}
	}
	if len(goalAmountMap) != 0 {
		for _, val := range goalAmountMap {
			totalSavings -= val.Amount
		}
	}

	return totalSavings
}