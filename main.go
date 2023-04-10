package main

import (
	"log"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/src/webapi"
)

func main() {

	// Initialize Database
	client, ctx, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Start Router
	webapi.StartRouter()


	// Add go routines here
		// create new bills & budgets here according to the time. check once every day

}
