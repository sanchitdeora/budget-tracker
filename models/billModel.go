package models

import (
	"log"
	"strings"
	"time"
)

type Bill struct {
	BillId 	  	    string 	  `json:"bill_id"`
	Title     	    string 	  `json:"title"`
	Category  	    string 	  `json:"category"`
	AmountDue 	    float32   `json:"amount_due"`
	DatePaid  	    time.Time `json:"date_paid"`
	DueDate   	    time.Time `json:"due_date"`
	Frequency  	    string 	  `json:"frequency"`
	Note   	  	    string 	  `json:"note,omitempty"`
	IsPaid 	  	    bool	  `json:"is_paid"`
	CreationTime    time.Time `json:"creation_time"`
	SequenceStartId string    `json:"sequence_start_id"`
	SequenceNumber  int       `json:"sequence_no"`
	Account		    string    `json:"account"`
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
	QUARTERLY_FREQUENCY,
	HALF_YEARLY_FREQUENCY,
	YEARLY_FREQUENCY,
}

func (bill *Bill) SetByUser() {
	if bill.CreationTime.IsZero() {
		bill.CreationTime = time.Now().Local()
	}
	bill.SequenceNumber = 0
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
	if bill.Frequency == "" {
		bill.Frequency = ONCE_FREQUENCY
		return
	}
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

func (bill *Bill) IsValid() bool {
	var invalidErr []string
	if bill.Title == "" {
		invalidErr = append(invalidErr, "title cannot be empty")
	}
	if bill.AmountDue < 0 {
		invalidErr = append(invalidErr, "amount cannot be less than zero")
	}

	// Account will be added later

	// if bill.Account != "" {
	// 	invalidErr = append(invalidErr, "account cannot be empty")
	// 	return false
	// }

	if len(invalidErr) > 0 {
		log.Println("Bill is invalid for the following reasons: ", strings.Join(invalidErr, ", "))
		return false
	}
	
	return true
}