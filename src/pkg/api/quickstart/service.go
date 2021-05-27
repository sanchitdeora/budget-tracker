package quickstart

import (
	"context"
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/models"
)


func SurveyService(ctx context.Context, survey models.Survey) error {
	err := db.AddSurvey(ctx, survey)
	if err != nil {
		log.Fatal(err)
	}
	return err
}