package main

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/pkg/bill"
	"github.com/sanchitdeora/budget-tracker/pkg/budget"
	"github.com/sanchitdeora/budget-tracker/pkg/goal"
	"github.com/sanchitdeora/budget-tracker/pkg/transaction"
	"github.com/sanchitdeora/budget-tracker/pkg/webapi"
)


func main() {

	// Initialize Database
	client, ctx, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)


	// database := db.NewDatabase()

	database := db.NewDatabase() 

	transactionService := transaction.NewService(&transaction.Opts{DB: database})
	billService := bill.NewService(&bill.Opts{TransactionService: transactionService, DB: database})
	goalService := goal.NewService(&goal.Opts{DB: database})

	budgetService := budget.NewService(&budget.Opts{
		TransactionService: transactionService,
		BillService: billService,
		GoalService: goalService,
		DB: database,
	})

	service := &webapi.ApiService{
		TransactionService: transactionService,
		BillService: billService,
		BudgetService: budgetService,
		GoalService: goalService,
	}

	// Add go routines here
	// create new bills & budgets here according to the time. check once every day
	go runMaintainer(service)

	// Start Router
	webapi.StartRouter(service)
}

func runMaintainer(s *webapi.ApiService) {
	dur, _ := time.ParseDuration("24h")
	updateBudgetTicker := time.NewTicker(dur)
	n := 0
	for {
		select {
			case <-updateBudgetTicker.C:
				go updateBills(s, n)
				updateBudgets(s, n)
				n++
			}
	}
}

func updateBudgets(s *webapi.ApiService, n int) {
	log.Println("running updateBudgets. No:", n)
	s.BudgetService.BudgetMaintainer(context.Background())
}

func updateBills(s *webapi.ApiService, n int) {
	log.Println("running updateBills. No:", n)
	s.BillService.BillMaintainer(context.Background())
}