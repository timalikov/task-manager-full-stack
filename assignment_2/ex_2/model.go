package ex_2

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}
