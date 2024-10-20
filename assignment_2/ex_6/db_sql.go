package ex_6

import (
	"assignment_2/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUsersSQL fetches users with optional filtering and pagination
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
// @Router /sql/users [get]
func GetUsersSQL(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ageQuery := c.Query("age")
		sortBy := c.DefaultQuery("sort_by", "name")
		pageQuery := c.DefaultQuery("page", "1")
		pageSizeQuery := c.DefaultQuery("page_size", "10")

		page, err := strconv.Atoi(pageQuery)
		if err != nil || page < 1 {
			sendValidationErrorResponse(c, "Invalid page parameter")
			return
		}

		pageSize, err := strconv.Atoi(pageSizeQuery)
		if err != nil || pageSize < 1 {
			sendValidationErrorResponse(c, "Invalid page_size parameter")
			return
		}

		offset := (page - 1) * pageSize

		query := "SELECT id, name, age FROM users WHERE 1=1"
		args := []interface{}{}

		if ageQuery != "" {
			age, err := strconv.Atoi(ageQuery)
			if err != nil {
				sendValidationErrorResponse(c, "Invalid age parameter")
				return
			}
			query += " AND age = $1"
			args = append(args, age)
		}

		query += fmt.Sprintf(" ORDER BY %s LIMIT $2 OFFSET $3", sortBy)
		args = append(args, pageSize, offset)

		rows, err := db.Query(query, args...)
		if err != nil {
			sendDatabaseErrorResponse(c, err)
			return
		}
		defer rows.Close()

		var users []model.User
		for rows.Next() {
			var user model.User
			if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
				sendDatabaseErrorResponse(c, err)
				return
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			sendDatabaseErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
