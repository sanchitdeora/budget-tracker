package models

type AccountResponse struct {
	User	UserResponse `json:"user"`
	Survey  SurveyResponse `json:"survey,omitempty"` 
}
