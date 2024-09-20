package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	ID      uint    `gorm:"primaryKey;autoIncrement"`
	Name    string  `gorm:"not null"`
	Age     int     `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	ID                uint   `gorm:"primaryKey;autoIncrement"`
	CustomerID        uint   `gorm:"not null;unique"`
	Bio               string `gorm:"type:text"`
	ProfilePictureURL string `gorm:"type:text"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Customer{}, &Profile{})
	if err != nil {
		return err
	}
	return nil
}
