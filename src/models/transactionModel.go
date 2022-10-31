package models

import (
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
	"auto_and_transport",
	"bills_and_utilities",
	"education",
	"entertainment",
	"food_and_dining",
	"health_and_fitness",
	"home",
	"income",
	"investments",
	"personal_care",
	"pets",
	"shopping",
	"taxes",
	"travel",
	"uncategorized",
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