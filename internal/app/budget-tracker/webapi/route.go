package webapi

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/pkg/api/home"
	"github.com/sanchitdeora/budget-tracker/pkg/api/quickstart"
	"github.com/sanchitdeora/budget-tracker/pkg/api/registration"
)

func StartRouter() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Routing
	route(router)

	// Start listener port
	router.Run()
}

func route(router *gin.Engine) *gin.Engine {

	router.POST("/quickstart", quickstart.OpeningSurvey)
	router.GET("/ping", abc)
	router.POST("/login", registration.Login)
	router.POST("/signup", registration.Signup)
	router.GET("/home", home.GetStarted)

	return router
}

func abc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
