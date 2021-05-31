package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/home"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/quickstart"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/registration"
)

func route(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")

	api.POST("/quickstart", quickstart.OpeningSurvey)
	api.POST("/login", registration.Login)
	api.POST("/register", registration.Register)
	api.GET("/home", home.GetStarted)
	api.GET("/ping", abc)

	return router
}

func abc(c *gin.Context) {
	c.JSON(201, gin.H{
		"message":"pong",
	})
}