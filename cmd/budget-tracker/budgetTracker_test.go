package budgetTracker_test

import (
	"testing"

	budgetTracker "github.com/sanchitdeora/budget-tracker/cmd/budget-tracker"
)

func TestBudgetting(t *testing.T) {
	if budgetTracker.Budgetting() != "Let's start budgeting!.." {
		t.Fatal("wronggg")
	}
}
