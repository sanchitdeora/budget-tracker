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
	Frequency  string 	`json:"frequency"`
	Note   	  string 	`json:"note,omitempty"`
	IsPaid 	  bool	 	`json:"is_paid"`
}

var BillCategoryMap = []string{
	BILLS_AND_UTILITIES_CATEGORY,
	EDUCATION_CATEGORY,
	ENTERTAINMENT_CATEGORY,
	LOANS_CATEGORY,
	MEDICAL_CATEGORY,
	RENT_CATEGORY,
	UNCATEGORIZED_CATEGORY,
}

var BillFrequencyMap = []string{
	ONCE_FREQUENCY,
	DAILY_FREQUENCY,
	WEEKLY_FREQUENCY,
	BI_WEEKLY_FREQUENCY,
	MONTHLY_FREQUENCY,
	BI_MONTHLY_FREQUENCY,
	QUATERLY_FREQUENCY,
	HALF_YEARLY_FREQUENCY,
	YEARLY_FREQUENCY,
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

func (bill *Bill) SetFrequency() {
	for index, frequency := range BillFrequencyMap {
		if frequency == strings.ToLower(bill.Frequency) {
			bill.Frequency = BillFrequencyMap[index]
			return
		}
		if index == len(BillFrequencyMap) - 1 {
			bill.Frequency = BillFrequencyMap[0]
		}
	}
}