package ex_3

import (
	"assignment_2/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsersGorm(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	}
}

func CreateUserGorm(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	}
}

func UpdateUserGorm(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Model(&model.User{}).Where("id = ?", id).Updates(user)
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUserGorm(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		db.Delete(&model.User{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
