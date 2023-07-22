package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/pkg/home"
	"github.com/sanchitdeora/budget-tracker/pkg/quickstart"
	"github.com/sanchitdeora/budget-tracker/pkg/registration"
)

func route(router *gin.Engine, service *ApiService) *gin.Engine {
	api := router.Group("/api")

	api.POST("/quickstart", quickstart.OpeningSurvey)
	api.POST("/login", registration.Login)
	api.POST("/register", registration.Register)
	api.GET("/home", home.GetStarted)
	api.GET("/ping", abc)

	// transactions
	api.GET("/transactions", service.GetAllTransactions)
	api.GET("/transaction/:id", service.GetTransactionById)
	api.POST("/transaction", service.CreateTransaction)
	api.PUT("/transaction/:id", service.UpdateTransactionById)
	api.DELETE("/transaction/:id", service.DeleteTransactionById)


	api.GET("/transactions/startTime/:startEpoch/endTime/:endEpoch", service.GetAllTransactionsByDate)

	// bills
	api.GET("/bills", service.GetAllBills)
	api.GET("/bill/:id", service.GetBillById)
	api.POST("/bill", service.CreateBill)
	api.PUT("/bill/:id", service.UpdateBillById)
	api.PUT("/bill/updateIsPaid/:id", service.UpdateBillIsPaid)
	api.PUT("/bill/updateIsUnpaid/:id", service.UpdateBillIsUnpaid)
	api.DELETE("/bill/:id", service.DeleteBillById)

	// budgets
	api.GET("/budgets", service.GetAllBudgets)
	api.GET("/budget/:id", service.GetBudgetById)
	api.POST("/budget", service.CreateBudget)
	api.PUT("/budget/:id", service.UpdateBudgetById)
	api.DELETE("/budget/:id", service.DeleteBudgetById)

	// goals
	api.GET("/goals", service.GetAllGoals)
	api.GET("/goal/:id", service.GetGoalById)
	api.POST("/goal", service.CreateGoal)
	api.PUT("/goal/:id", service.UpdateGoalById)
	api.DELETE("/goal/:id", service.DeleteGoalById)

	return router
}

func abc(c *gin.Context) {
	c.JSON(201, gin.H{
		"message":"pong",
	})
}