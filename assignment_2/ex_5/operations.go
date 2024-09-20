package ex_5

import (
	"assignment_2/model"
	"fmt"
	"gorm.io/gorm"
)

func InsertCustomerWithProfile(db *gorm.DB, customer *model.Customer, profile *model.Profile) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(customer).Error; err != nil {
			return err
		}
		profile.CustomerID = customer.ID
		if err := tx.Create(profile).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to insert customer and profile: %v", err)
	}
	return nil
}

func QueryCustomerWithProfile(db *gorm.DB, customerId uint) (*model.Customer, error) {
	var customer model.Customer
	if err := db.Preload("Profile").First(&customer, customerId).Error; err != nil {
		return nil, fmt.Errorf("failed to query customer with profile: %v", err)
	}
	return &customer, nil
}

func UpdateCustomerProfile(db *gorm.DB, customerId uint, bio string, profilePicURL string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var profile model.Profile
		if err := tx.Where("customer_id = ?", customerId).First(&profile).Error; err != nil {
			return err
		}
		profile.Bio = bio
		profile.ProfilePictureURL = profilePicURL
		if err := tx.Save(&profile).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update profile: %v", err)
	}
	return nil
}

func DeleteCustomerWithProfile(db *gorm.DB, customerId uint) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Customer{}, customerId).Error; err != nil {
			return err
		}
		if err := tx.Where("customer_id = ?", customerId).Delete(&model.Profile{}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete customer and profile: %v", err)
	}
	return nil
}
