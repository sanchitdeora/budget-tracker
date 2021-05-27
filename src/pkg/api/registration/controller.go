package registration

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/models"
)

func Signup(c *gin.Context) {

	userBody := new(models.User)
	err := c.BindJSON(&userBody)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
	}

	// Store in DB
	err = SignupService(c, *userBody)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Signup",
		"body":    userBody,
	})
}

func Login(c *gin.Context) {

	loginBody := new(models.Login)
	err := c.BindJSON(&loginBody)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failure",
			"error":   err,
		})
	}

	// Check if login is correct
	err = LoginService(c, *loginBody)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "validation error",
			"error":   err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Login",
	})
}
