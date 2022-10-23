package models

import (
	"strings"
)

type Transaction struct {
	TransactionId string `json:"transactionId"`
	Title         string `json:"title"`
	Category      string `json:"category"`
	Amount        int    `json:"amount"`
	Date          int    `json:"date"`
	Description   string `json:"description"`
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
		if category != strings.ToLower(transaction.Category) && index == len(TransactionCategoryMap)-1 {
			transaction.Category = TransactionCategoryMap[index]
		}
	}
}