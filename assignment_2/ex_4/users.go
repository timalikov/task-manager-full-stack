package ex_4

import (
	"assignment_2/model"
	"fmt"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User2{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %v", err)
	}
	fmt.Println("Users table created successfully using GORM AutoMigrate!")
	return nil
}

func InsertUsersWithTransaction(db *gorm.DB, users []model.User2) error {
	tx := db.Begin()

	for _, user := range users {
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert user %v: %v", user.Name, err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	fmt.Println("Users inserted successfully using GORM transactions!")
	return nil
}

func QueryUsersWithPagination(db *gorm.DB, age int, page int, pageSize int) ([]model.User2, error) {
	var users []model.User2
	offset := (page - 1) * pageSize

	query := db.Offset(offset).Limit(pageSize)
	if age > 0 {
		query = query.Where("age = ?", age)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query users with GORM: %v", err)
	}

	return users, nil
}

func UpdateUser(db *gorm.DB, id int, name string, age int) error {
	err := db.Model(&model.User2{}).Where("id = ?", id).Updates(model.User2{Name: name, Age: age}).Error
	if err != nil {
		return fmt.Errorf("failed to update user with GORM: %v", err)
	}
	fmt.Printf("User with ID %d updated successfully using GORM!\n", id)
	return nil
}

func DeleteUser(db *gorm.DB, id int) error {
	err := db.Delete(&model.User2{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %d using GORM: %v", id, err)
	}
	fmt.Printf("User with ID %d deleted successfully using GORM!\n", id)
	return nil
}
