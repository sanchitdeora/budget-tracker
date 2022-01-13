package account

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserAccount(c *gin.Context) {

	userID := c.Param("userId")

	// Check if login is correct
	resp, err := GetAccount(c, userID)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "validation error",
			"error":   err,
		})
	}

	c.JSON(200, resp)
}
