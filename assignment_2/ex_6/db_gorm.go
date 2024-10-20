package ex_6

import (
	"assignment_2/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetUsersGORM fetches users with optional filtering and pagination
// @Summary Get all users
// @Description Get users with optional filtering by age and pagination
// @Tags users
// @Accept  json
// @Produce  json
// @Param   age       query   int     false  "Filter by Age"
// @Param   page      query   int     false  "Page number"
// @Param   page_size query   int     false  "Page size"
// @Param   sort_by   query   string  false  "Sort by column name"
// @Success 200 {array} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /gorm/users [get]
func GetUsersGORM(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []model.User
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
