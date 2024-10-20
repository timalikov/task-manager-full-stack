package migrations

import (
	"log"
	"task_management_backend/db"
	"task_management_backend/models"
)

func Migrate() {
	database, err := db.ConnectToDb()
	if err != nil {
		log.Fatalf("failed to connect to the database for migration: %v", err)
	}

	err = database.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Database migration completed successfully!")
}
