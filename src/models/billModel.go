package models

import (
	"strings"
)

type Bill struct {
	BillId 	  string `json:"bill_id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	AmountDue int    `json:"amount_due"`
	DueDate   int    `json:"due_date"`
	HowOften  string `json:"how_often"`
	Note   	  string `json:"note,omitempty"`
	IsPaid 	  bool	 `json:"is_paid"`
}

var BillCategoryMap = []string{
	"bills_and_utilities",
	"rent",
	"medical",
	"education",
	"loans",
	"day_care",
	"uncategorized",
}

var BillHowOftenMap = []string{
	"once",
	"weekly",
	"bi-weekly",
	"monthly",
	"bi-monthly",
	"quaterly",
	"quaterly",
	"half_yearly",
	"yearly",
}

func (bill *Bill) SetCategory() {
	for index, category := range BillCategoryMap {
		if category == strings.ToLower(bill.Category) {
			bill.Category = BillCategoryMap[index]
			return
		}
		if index == len(BillCategoryMap)-1 {
			bill.Category = BillCategoryMap[index]
		}
	}
}

func (bill *Bill) SetHowOften() {
	for index, howOften := range BillHowOftenMap {
		if howOften == strings.ToLower(bill.HowOften) {
			bill.HowOften = BillHowOftenMap[index]
			return
		}
		if index == len(BillHowOftenMap) - 1 {
			bill.HowOften = BillHowOftenMap[0]
		}
	}
}