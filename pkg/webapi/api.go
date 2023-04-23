package webapi

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sanchitdeora/budget-tracker/pkg/bill"
	"github.com/sanchitdeora/budget-tracker/pkg/budget"
	"github.com/sanchitdeora/budget-tracker/pkg/goal"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
)

func StartRouter(service *ApiService) {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// Routing
	route(router, service)

	// Start listener port
	router.Run()
}

type ApiService	struct {
	TransactionService 	transaction.Service
	BillService 		bill.Service
	BudgetService 		budget.Service
	GoalService 		goal.Service
}
