package main

import (
	"context"

	"github.com/Hertucktor/authapi/api"
	"github.com/Hertucktor/authapi/dbhandler"
)

func main() {
	// connect to MongoDB
	dbhandler.InitDB()
	defer dbhandler.Client.Disconnect(context.TODO())

	api.SetupRoutes()
}
