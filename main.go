package main

import (
	"fmt"
	"log"

	"github.com/sanchitdeora/budget-tracker/cmd/budget-tracker"
	"github.com/sanchitdeora/budget-tracker/internal/app/budget-tracker/webapi"
	"github.com/sanchitdeora/budget-tracker/repository"
)

func main() {
	fmt.Println(budgetTracker.Budgetting())

	// Initialize Database
	client, ctx, err := repository.Init()
	if err != nil {
		log.Fatal(err)
	}
	
	defer client.Disconnect(ctx)

	// Start Router
	webapi.StartRouter()
	
}