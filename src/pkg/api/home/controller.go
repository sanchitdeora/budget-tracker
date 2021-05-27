package home

import "github.com/gin-gonic/gin"

func GetStarted(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Success",
	})
}

