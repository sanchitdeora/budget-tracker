package quickstart

import (
	"context"
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
)

func SurveyService(ctx context.Context, survey models.Survey) error {
	user, err := db.GetUserByEmail(ctx, survey.Email)
	if err != nil {
		log.Println(err)
		return err
	}

	survey.UserID = user.UserID
	err = db.AddSurvey(ctx, &survey)
	if err != nil {
		log.Println(err)
		return err
	}
	user.SurveyID = survey.SurveyID
	user.SurveyComplete = true
	err = db.UpdateUser(ctx, *user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
