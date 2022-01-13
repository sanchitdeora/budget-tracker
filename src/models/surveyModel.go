package models

type Survey struct {
	SurveyID      string `json:"surveyId"`
	UserID        string `json:"userId"`
	Email         string `json:"email"`
	MonthlyIncome string  `json:"monthlyIncome"`
	SavingsType   string `json:"savingsType"`
	MonthlyLimit  string  `json:"monthlyLimit"`
}

type SurveyResponse struct {
	MonthlyIncome string  `json:"monthlyIncome"`
	SavingsType   string `json:"savingsType"`
	MonthlyLimit  string  `json:"monthlyLimit"`
}
