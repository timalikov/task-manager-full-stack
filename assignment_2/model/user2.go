package model

type User2 struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
	Age  int    `gorm:"not null" json:"age"`
}
