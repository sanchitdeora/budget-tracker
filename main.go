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

}
