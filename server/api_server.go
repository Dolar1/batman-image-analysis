package server

import (
	"batman/app"
	"batman/routes"
	"log"
)

func StartAPIServer(dependency *app.Dependency) {
	router := routes.SetupRoutes(dependency)
	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
