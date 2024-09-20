package ex_2

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=timalikov password=1234 dbname=go_assignment2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	log.Println("Connected to the database successfully!")
	return db
}
