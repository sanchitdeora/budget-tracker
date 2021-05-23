package registration

import (
	"context"

	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/repository"
)


func SignupService(c context.Context, user models.User) error {
	repository.CreateAccountRecords(c, user)
	return nil
}