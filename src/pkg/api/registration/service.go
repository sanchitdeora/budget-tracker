package registration

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func RegisterService(ctx context.Context, user models.User) error {
	err := db.AddUser(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func LoginService(ctx context.Context, login models.Login) error {
	// improve validation logic

	loginFromDB, err := db.GetLoginInfo(ctx, login)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(loginFromDB)
	fmt.Println(*loginFromDB)

	if loginFromDB.Password != login.Password {
		return errors.New("invaid login")
	}

	return err
}
