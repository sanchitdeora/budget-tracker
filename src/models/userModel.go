package models

type User struct {
	UserID         string `json:"userId" bson:"userId"`
	Firstname      string `json:"firstname" bson:"firstname"`
	Lastname       string `json:"lastname" bson:"lastname"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
	DateOfBirth    string `json:"dateOfBirth" bson:"dateOfBirth"`
	PhoneNumber    string `json:"phoneNumber" bson:"phoneNumber"`
	SurveyID       string `json:"surveyId,omitempty" bson:"surveyId,omitempty"`
	SurveyComplete bool   `json:"surveyComplete" bson:"surveyComplete"`
}

type UserResponse struct {
	UserID      string `json:"userId"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
	PhoneNumber string `json:"phoneNumber"`
}
