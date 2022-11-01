package models

import (
	"strings"
	"time"
)

type Bill struct {
	BillId 	  string 	`json:"bill_id"`
	Title     string 	`json:"title"`
	Category  string 	`json:"category"`
	AmountDue float32   `json:"amount_due"`
	DatePaid  time.Time `json:"date_paid"`
	DueDate   time.Time `json:"due_date"`
	HowOften  string 	`json:"how_often"`
	Note   	  string 	`json:"note,omitempty"`
	IsPaid 	  bool	 	`json:"is_paid"`
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
	"bi_weekly",
	"monthly",
	"bi_monthly",
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