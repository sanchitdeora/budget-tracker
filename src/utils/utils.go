package utils

import (
	"errors"
	"time"

	"github.com/sanchitdeora/budget-tracker/src/models"
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

func CalculateEndDateWithFrequency(currDate time.Time, freq string) (time.Time, error) {
	err := errors.New("provided Frequency is not found in the frequency map")

	if !Contains(models.BillFrequencyMap, freq) || models.ONCE_FREQUENCY == freq {
		return currDate, err
	}
	if freq == models.DAILY_FREQUENCY {
		return currDate.AddDate(0, 0, 1), nil
	} else if freq == models.WEEKLY_FREQUENCY {
		return currDate.AddDate(0, 0, 7), nil
	} else if freq == models.BI_WEEKLY_FREQUENCY {
		return currDate.AddDate(0, 0, 14), nil
	} else if freq == models.MONTHLY_FREQUENCY {
		return currDate.AddDate(0, 1, 0), nil
	} else if freq == models.BI_MONTHLY_FREQUENCY {
		return currDate.AddDate(0, 2, 0), nil
	} else if freq == models.QUARTERLY_FREQUENCY {
		return currDate.AddDate(0, 3, 0), nil
	} else if freq == models.HALF_YEARLY_FREQUENCY {
		return currDate.AddDate(0, 6, 0), nil
	} else if freq == models.YEARLY_FREQUENCY {
		return currDate.AddDate(1, 0, 0), nil
	} else {
		return currDate, err
	}

}