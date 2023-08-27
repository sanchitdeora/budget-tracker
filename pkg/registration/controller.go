package registration

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)


func Register(c *gin.Context) {

	userBody := new(models.User)
	err := c.BindJSON(&userBody)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
	}

	// Store in DB
	err = RegisterService(c, *userBody)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Register",
		"body":    userBody,
	})
}

func Login(c *gin.Context) {

	loginBody := new(models.Login)
	err := c.BindJSON(&loginBody)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
	}

	// Check if login is correct
	err = LoginService(c, loginBody)

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "validation error",
			"error":   err,
		})
	}

	log.Println("LoginBody outside:", loginBody)

	c.JSON(200, gin.H{
		"message":          "Login",
		"token":            "token123",
		"isSurveyComplete": loginBody.SurveyComplete,
	})
}
