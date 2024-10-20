package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strconv"
	"task_management_backend/config"
)

type BasePsqlDTO struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Conn     *gorm.DB
}

func ConnectToDb() (*gorm.DB, error) {
	cfg := config.LoadConfig()

	dbPort, err := strconv.Atoi(cfg.DbPort)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.Host, dbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	log.Println("Successfully connected to the database using GORM")

	return db, nil
}
