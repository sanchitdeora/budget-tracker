package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message        string `json:"msg"`
	Token          string `json:"token"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	SurveyComplete bool   `json:"surveyComplete"`
}
