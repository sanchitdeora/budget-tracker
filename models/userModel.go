package models

type User struct {
	FirstName 	string	`json:"firstName"`
	LastName 	string	`json:"lastName"`
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
	DateOfBirth string	`json:"dateOfBirth"`
	PhoneNumber string 	`json:"phoneNumber"`
}

type Login struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}