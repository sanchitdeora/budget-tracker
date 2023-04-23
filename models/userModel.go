package models

type User struct {
	UserID         string `json:"userId"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	DateOfBirth    string `json:"dateOfBirth"`
	PhoneNumber    string `json:"phoneNumber"`
	SurveyID       string `json:"surveyId,omitempty"`
	SurveyComplete bool   `json:"surveyComplete"`
}

type Login struct {
	Email    		string `json:"email"`
	Password 		string `json:"password"`
	SurveyComplete 	bool   `json:"surveyComplete"`
}
