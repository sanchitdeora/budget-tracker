package models

type Survey struct {
	UserID        int64  `json:"user_id,omitempty"`
	MonthlyIncome int64  `json:"monthlyIncome"`
	SavingsType   string `json:"savingsType"`
	MonthlyLimit  int64  `json:"monthlyLimit"`
}
