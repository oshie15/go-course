package main

import (
	"context"
	"net/http"

	"github.com/oshie15/go-course.git/config"
	"github.com/oshie15/go-course.git/db"
	"github.com/oshie15/go-course.git/routes"
)

func main() {
	// Load routes
	handler := routes.MounteRoutes()

	// Connect to DB
	db.InitDB()
	defer db.DB.Close(context.Background())

	// Start server
	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}
	server.ListenAndServe()
}
