package quickstart

import "github.com/gin-gonic/gin"

type Survey struct {
	MonthlyIncome int64  `json:"monthlyIncome"`
	SavingsType   string `json:"savingsType"`
	MonthlyLimit  int64  `json:"monthlyLimit"`
}

func OpeningSurvey(c *gin.Context) {

	reqBody := new(Survey)
	err := c.BindJSON(reqBody)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error": err,
		})
		panic(err)
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"body":    reqBody,
	})
}
