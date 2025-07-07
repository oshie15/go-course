package main

import (
	"context"
	"net/http"

	"github.com/oshie15/go-course.git/config"
	"github.com/oshie15/go-course.git/db"
	"github.com/oshie15/go-course.git/routes"
)

func main() {
	handler := routes.MountRoutes()
	
	

	db.InitDB()
	server := &http.Server{
		Addr: config.Config.AppPort,
		Handler: handler,
	}
	defer db.DB.Close(context.Background())
	server.ListenAndServe()

}