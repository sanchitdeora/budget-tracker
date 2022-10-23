package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/bill"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/home"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/quickstart"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/registration"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/transaction"
)

func route(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")

	api.POST("/quickstart", quickstart.OpeningSurvey)
	api.POST("/login", registration.Login)
	api.POST("/register", registration.Register)
	api.GET("/home", home.GetStarted)
	api.GET("/ping", abc)

	api.GET("/transactions", transaction.GetAllTransactions)
	api.GET("/transaction/:id", transaction.GetTransactionById)
	api.POST("/transaction", transaction.CreateTransaction)
	api.PUT("/transaction/:id", transaction.UpdateTransaction)
	api.DELETE("/transaction/:id", transaction.DeleteTransaction)

	api.GET("/bills", bill.GetAllBills)
	api.GET("/bill/:id", bill.GetBillById)
	api.POST("/bill", bill.CreateBill)
	api.PUT("/bill/:id", bill.UpdateBill)
	api.DELETE("/bill/:id", bill.DeleteBill)

	return router
}

func abc(c *gin.Context) {
	c.JSON(201, gin.H{
		"message":"pong",
	})
}