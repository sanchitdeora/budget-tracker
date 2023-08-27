package utils

import (
	"time"

	"github.com/sanchitdeora/budget-tracker/models"
)

func Contains(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

func SearchIndex[T comparable](s []T, v T) int {
	for i, val := range s {
		if val == v {
			return i
		}
	}
	return -1
}

func Remove[T comparable](s []T, index int) []T{
    return append(s[:index], s[index+1:]...)
}

func CalculateEndDateWithFrequency(currDate time.Time, freq string) time.Time {
	if freq == models.DAILY_FREQUENCY {
		return currDate.AddDate(0, 0, 1)
	} else if freq == models.WEEKLY_FREQUENCY {
		return currDate.AddDate(0, 0, 7)
	} else if freq == models.BI_WEEKLY_FREQUENCY {
		return currDate.AddDate(0, 0, 14)
	} else if freq == models.MONTHLY_FREQUENCY {
		return currDate.AddDate(0, 1, 0)
	} else if freq == models.BI_MONTHLY_FREQUENCY {
		return currDate.AddDate(0, 2, 0)
	} else if freq == models.QUARTERLY_FREQUENCY {
		return currDate.AddDate(0, 3, 0)
	} else if freq == models.HALF_YEARLY_FREQUENCY {
		return currDate.AddDate(0, 6, 0)
	} else if freq == models.YEARLY_FREQUENCY {
		return currDate.AddDate(1, 0, 0)
	}

	return currDate
}