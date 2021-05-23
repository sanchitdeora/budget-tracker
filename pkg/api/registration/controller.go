package registration

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/models"
)

func Signup(c *gin.Context) {

	userBody := new(models.User)
	err := c.BindJSON(userBody)

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
	// Check if login is correct

	c.JSON(200, gin.H{
		"message": "Login",
	})
}
