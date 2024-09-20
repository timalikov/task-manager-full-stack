package ex_5

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=timalikov password=1234 dbname=go_assignment2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get db from GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to PostgreSQL with GORM!")
	return db
}
