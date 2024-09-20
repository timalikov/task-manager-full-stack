package ex_2

import (
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	log.Println("Table created successfully!")
}

func InsertData(db *gorm.DB) {
	users := []User{
		{Name: "Maks", Age: 25},
		{Name: "Alihan", Age: 30},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatalf("failed to insert user: %v", result.Error)
		}
		log.Printf("Inserted user with ID: %d\n", user.ID)
	}
}

func QueryData(db *gorm.DB) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatalf("failed to query users: %v", result.Error)
	}

	for _, user := range users {
		log.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
