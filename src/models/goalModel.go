package models

import (
	"time"
)

type Goal struct {
	GoalId			string		`json:"goal_id"`
	Name			string		`json:"name"`
	CurrentAmount	float32		`json:"current_amount"`
	TargetAmount	float32		`json:"target_amount"`
	TargetDate		time.Time	`json:"target_date"`
	BudgetIdList 	[]string	`json:"budget_id_list"`
}

// func (goal *Goal) GetAmountSavingPerBudgetsFrequency() (map[string]float32, error) {
// 	var amountMap map[string]float32
// 	for _, id := range goal.BudgetIdList {
// 		_, err := budget.GetBudget(context.Background(), id)
// 		if err != nil {
// 			return nil, errors.New("Unable to fetch Get Budget")
// 		}
// 		// amountMap[id] = calculateAmountSavingPerFrequency(goal, &response.Frequency)
// 	}
// 	return amountMap, nil
// }

// func assertValidBudgetFrequency(freq string) error {
// 	if !utils.Contains(BudgetFrequencyMap, freq) {
// 		return errors.New("provided Frequency is not found in the budget frequency map")
// 	}
// 	return nil
// }

// func calculateAmountSavingPerFrequency(goal *Goal, freq string) float32 {
// 	differenceDate := goal.TargetDate.Sub(time.Now())
// 	differenceDays := differenceDate.Hours() / 24

// 	balance := goal.TargetAmount - goal.CurrentAmount

// 	if freq == WEEKLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / 7)

// 	} else if freq == BI_WEEKLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / (7 * 2))

// 	} else if freq == MONTHLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / 30)

// 	} else if freq == BI_MONTHLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / (30 * 2))

// 	} else if freq == QUATERLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / (30 * 3))

// 	} else if freq == HALF_YEARLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / (30 * 6))

// 	} else if freq == YEARLY_FREQUENCY {
// 		return balance / (float32(differenceDays) / 365)

// 	}

// 	return balance
// }