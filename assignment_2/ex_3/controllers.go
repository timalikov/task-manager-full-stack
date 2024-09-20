package ex_3

import (
	"assignment_2/ex_1"
	"assignment_2/ex_2"
	"assignment_2/model"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	dbGorm := ex_2.ConnectDB()

	dbDirect, err := ex_1.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database (SQL): %v", err)
	}

	dbGorm.AutoMigrate(&model.User{})

	r.GET("/gorm/users", GetUsersGorm(dbGorm))
	r.POST("/gorm/user", CreateUserGorm(dbGorm))
	r.PUT("/gorm/user/:id", UpdateUserGorm(dbGorm))
	r.DELETE("/gorm/user/:id", DeleteUserGorm(dbGorm))

	// Direct SQL routes
	r.GET("/direct/users", GetUsersDirect(dbDirect))
	r.POST("/direct/user", CreateUserDirect(dbDirect))
	r.PUT("/direct/user/:id", UpdateUserDirect(dbDirect))
	r.DELETE("/direct/user/:id", DeleteUserDirect(dbDirect))

	return r
}
