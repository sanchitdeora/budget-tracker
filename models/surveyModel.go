package models

type Survey struct {
	SurveyID      string `json:"surveyId"`
	UserID        string `json:"userId"`
	Email         string `json:"email"`
	MonthlyIncome int64  `json:"monthlyIncome"`
	SavingsType   string `json:"savingsType"`
	MonthlyLimit  int64  `json:"monthlyLimit"`
}
