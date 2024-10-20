package main

import (
	"log"
	"task_management_backend/routes"
)

func main() {

	// Migrations: it is enough to run it one time
	//migrations.Migrate()

	router := routes.SetupRoutes()

	// Start the server
	log.Println("Starting server on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
