package models

import (
	"strings"
)

type Transaction struct {
	TransactionId string `json:"transaction_id"`
	Title         string `json:"title"`
	Category      string `json:"category"`
	Amount        int    `json:"amount"`
	Date          int    `json:"date"`
	Note   string `json:"Note"`
}

var TransactionCategoryMap = []string{
	"auto_and_transport",
	"bills_and_utilities",
	"education",
	"entertainment",
	"food",
	"health",
	"income",
	"personal_care",
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