package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/bill"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/budget"
	"github.com/sanchitdeora/budget-tracker/src/pkg/api/goals"
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
	api.PUT("/bill/updateIsPaid/:id", bill.UpdateBillIsPaid)
	api.PUT("/bill/updateIsUnpaid/:id", bill.UpdateBillIsUnpaid)
	api.DELETE("/bill/:id", bill.DeleteBill)

	api.GET("/budgets", budget.GetAllBudgets)
	api.GET("/budget/:id", budget.GetBudgetById)
	api.POST("/budget", budget.CreateBudget)
	api.PUT("/budget/:id", budget.UpdateBudget)
	api.DELETE("/budget/:id", budget.DeleteBudget)

	api.GET("/goals", goals.GetAllGoals)
	api.GET("/goal/:id", goals.GetGoalById)
	api.POST("/goal", goals.CreateGoal)
	api.PUT("/goal/:id", goals.UpdateGoal)
	api.DELETE("/goal/:id", goals.DeleteGoal)

	return router
}

func abc(c *gin.Context) {
	c.JSON(201, gin.H{
		"message":"pong",
	})
}