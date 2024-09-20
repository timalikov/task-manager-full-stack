package ex_6

import (
	"assignment_2/ex_3"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutesSQL(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/sql/users", GetUsersSQL(db))
	r.POST("/sql/users", ex_3.CreateUserDirect(db))
	r.PUT("/sql/users/:id", ex_3.UpdateUserDirect(db))
	r.DELETE("/sql/users/:id", ex_3.DeleteUserDirect(db))

	return r
}

func SetupRoutesGORM(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/gorm/users", GetUsersGORM(db))
	r.POST("/gorm/users", ex_3.CreateUserGorm(db))
	r.PUT("/gorm/users/:id", ex_3.UpdateUserGorm(db))
	r.DELETE("/gorm/users/:id", ex_3.DeleteUserGorm(db))

	return r
}
