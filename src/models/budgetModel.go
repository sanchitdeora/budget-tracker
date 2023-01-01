package models

import (
	"strings"
)

type Budget struct {
	BudgetId         string             `json:"budget_id"`
	Name             string             `json:"name"`
	IncomeMap        map[string]float32 `json:"income_map"`
	SpendingLimitMap map[string]float32 `json:"spending_limit_map"`
	GoalAmountMap    map[string]float32 `json:"goal_amount_map"`
	Frequency        string             `json:"frequency"`
	Savings          float32            `json:"savings"`
}

var BudgetFrequencyMap = []string{
	ONCE_FREQUENCY,
	DAILY_FREQUENCY,
	WEEKLY_FREQUENCY,
	BI_WEEKLY_FREQUENCY,
	MONTHLY_FREQUENCY,
	BI_MONTHLY_FREQUENCY,
	QUATERLY_FREQUENCY,
	HALF_YEARLY_FREQUENCY,
	YEARLY_FREQUENCY,
}

func (budget *Budget) GetSavings() float32 {
	if budget.Savings == 0 {
		budget.Savings = calculateSavings(budget.IncomeMap, budget.SpendingLimitMap, budget.GoalAmountMap)
	}

	return budget.Savings
}

func (budget *Budget) SetFrequency() {
	for index, frequency := range BudgetFrequencyMap {
		if frequency == strings.ToLower(budget.Frequency) {
			budget.Frequency = BudgetFrequencyMap[index]
			return
		}
		// default frequency set as 
		if index == len(BudgetFrequencyMap) - 1 {
			budget.Frequency = ONCE_FREQUENCY
		}
	}
}

func (budget *Budget) SetCategory() {
	// replace categories with well defined values in map
	// replace unknown categories with uncategorized
	for spendCategory, limit := range budget.SpendingLimitMap {
		for index, categoryFromMap := range TransactionCategoryMap {
			if categoryFromMap == strings.ToLower(spendCategory) {
				delete(budget.SpendingLimitMap, spendCategory)
				budget.SpendingLimitMap[categoryFromMap] = limit

				break
			}
			if index == len(TransactionCategoryMap)-1 {
				delete(budget.SpendingLimitMap, spendCategory)
				if budget.SpendingLimitMap[TransactionCategoryMap[index]] > 0 {
					budget.SpendingLimitMap[TransactionCategoryMap[index]] = budget.SpendingLimitMap[TransactionCategoryMap[index]] + limit	
				} else {
					budget.SpendingLimitMap[TransactionCategoryMap[index]] = limit
				}
			}
		}
	}
}

func calculateSavings(incomeMap map[string]float32, spendingLimitMap map[string]float32, goalAmountMap map[string]float32) float32 {
	var totalSavings float32
	for _, val := range incomeMap {
		totalSavings += val
	}

	if len(spendingLimitMap) != 0 {
		for _, val := range spendingLimitMap {
			totalSavings -= val
		}
	}
	if len(goalAmountMap) != 0 {
		for _, val := range goalAmountMap {
			totalSavings -= val
		}
	}

	return totalSavings
}