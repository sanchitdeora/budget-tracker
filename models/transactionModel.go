package models

import (
	"log"
	"strings"
	"time"
)

type Transaction struct {
	TransactionId string 	`json:"transaction_id"`
	Title         string 	`json:"title"`
	Category      string 	`json:"category"`
	Amount        float32   `json:"amount"`
	Date          time.Time `json:"date"`
	Type		  bool		`json:"type"` // true if Credit, false if Debit
	Account		  string    `json:"account"`
	Note   		  string 	`json:"note"`
}

var TransactionCategoryMap = []string{
	AUTO_AND_TRANSPORT_CATEGORY,
	BILLS_AND_UTILITIES_CATEGORY,
	EDUCATION_CATEGORY,
	ENTERTAINMENT_CATEGORY,
	FOOD_AND_DINING_CATEGORY,
	HEALTH_AND_FITNESS_CATEGORY,
	HOME_CATEGORY,
	INCOME_CATEGORY,
	INVESTMENTS_CATEGORY,
	PERSONAL_CARE_CATEGORY,
	PETS_CATEGORY,
	SHOPPING_CATEGORY,
	TAXES_CATEGORY,
	TRAVEL_CATEGORY,
	UNCATEGORIZED_CATEGORY,
}

func (transaction *Transaction) SetCategory() {
	for index, category := range TransactionCategoryMap {
		if category == strings.ToLower(transaction.Category) {
			transaction.Category = TransactionCategoryMap[index]
			return
		}
		if index == len(TransactionCategoryMap)-1 {
			transaction.Category = TransactionCategoryMap[index]
		}
	}
}

func (transaction *Transaction) FromBill(bill *Bill) {
	transaction.Title = bill.Title
	transaction.Category = BILLS_AND_UTILITIES_CATEGORY
	transaction.Amount = bill.AmountDue
	transaction.Date = bill.DatePaid
	transaction.Type = false
	transaction.Account = ""
	transaction.Note = bill.Note

}

func (transaction *Transaction) IsValid() bool {
	var invalidErr []string
	if transaction.Title == "" {
		invalidErr = append(invalidErr, "title cannot be empty")
	}
	if transaction.Amount < 0 {
		invalidErr = append(invalidErr, "amount cannot be less than zero")
	}

	// Account will be added later

	// if transaction.Account != "" {
	// 	invalidErr = append(invalidErr, "account cannot be empty")
	// 	return false
	// }

	if len(invalidErr) > 0 {
		log.Println("transaction is invalid for the following reasons:", strings.Join(invalidErr, ", "))
		return false
	}

	return true
}