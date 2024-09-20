package ex_6

import (
	"assignment_2/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetUsersGORM(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []model.User2
		ageQuery := c.Query("age")
		sortBy := c.DefaultQuery("sort_by", "name")
		pageQuery := c.DefaultQuery("page", "1")
		pageSizeQuery := c.DefaultQuery("page_size", "10")

		page, _ := strconv.Atoi(pageQuery)
		pageSize, _ := strconv.Atoi(pageSizeQuery)
		offset := (page - 1) * pageSize

		query := db.Offset(offset).Limit(pageSize).Order(sortBy)

		if ageQuery != "" {
			query = query.Where("age = ?", ageQuery)
		}

		if err := query.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
