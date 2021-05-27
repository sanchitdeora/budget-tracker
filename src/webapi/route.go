package webapi

import (

	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/home"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/quickstart"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/registration"
)

func route(router *gin.Engine) *gin.Engine {

	router.POST("/quickstart", quickstart.OpeningSurvey)
	router.POST("/login", registration.Login)
	router.POST("/signup", registration.Signup)
	router.GET("/home", home.GetStarted)

	return router
}
