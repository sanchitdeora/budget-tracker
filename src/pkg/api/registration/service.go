package registration

import (
	"context"
	"crypto/rand"
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

func LoginService(ctx context.Context, login *models.Login) (*models.LoginResponse, error) {
	// improve validation logic

	user, err := db.GetUserByEmail(ctx, login.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if user.Password != login.Password {
		return nil, errors.New("invaid login")
	}

	var resp models.LoginResponse

	resp.UserID = user.UserID
	resp.Message = "login response"
	resp.Token = tokenGenerator()
	resp.Email = login.Email
	resp.Name = user.Lastname + ", " + user.Firstname
	resp.SurveyComplete = user.SurveyComplete

	return &resp, err
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
