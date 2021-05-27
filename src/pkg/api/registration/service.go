package registration

import (
	"context"
	"errors"
	"log"

	"github.com/sanchitdeora/budget-tracker/src/models"
	"github.com/sanchitdeora/budget-tracker/db"
)

func SignupService(ctx context.Context, user models.User) error {
	err := db.AddSignup(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func LoginService(ctx context.Context, login models.Login) error {
	// improve validation logic
	
	loginFromDB, err := db.GetLoginInfo(ctx, login)
	if err != nil {
		log.Fatal(err)
	}

	if loginFromDB != login {
		return errors.New("invaid login")
	}

	return err
}
