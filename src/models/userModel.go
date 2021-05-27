package models

type User struct {
	UserID		int64	`json:"user_id"`
	Firstname 	string	`json:"firstname"`
	Lastname 	string	`json:"lastname"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
	DateOfBirth string	`json:"dateOfBirth"`
	PhoneNumber string 	`json:"phoneNumber"`
}

type Login struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}