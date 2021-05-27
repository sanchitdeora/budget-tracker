package quickstart

import (
	"context"
	"log"

	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/repository"
)


func SurveyService(ctx context.Context, survey models.Survey) error {
	err := repository.AddSurvey(ctx, survey)
	if err != nil {
		log.Fatal(err)
	}
	return err
}