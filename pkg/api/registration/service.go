package registration

import (
	"context"
	"errors"
	"log"

	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/repository"
)

func SignupService(c context.Context, user models.User) error {
	err := repository.CreateUserRecord(c, user)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func LoginService(c context.Context, login models.Login) error {
	loginFromDB, err := repository.GetLoginInfo(c, login)
	if err != nil {
		log.Fatal(err)
	}

	if loginFromDB != login {
		return errors.New("invaid login")
	}

	return nil
}
