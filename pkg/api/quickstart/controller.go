package quickstart

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)

func OpeningSurvey(c *gin.Context) {

	reqBody := new(models.Survey)
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
		log.Fatal(err)
	}
	SurveyService(c, *reqBody)

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    reqBody,
	})
}
