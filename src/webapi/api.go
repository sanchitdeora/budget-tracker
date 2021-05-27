package webapi

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin")


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