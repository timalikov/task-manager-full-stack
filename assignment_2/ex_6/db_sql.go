package ex_6

import (
	"assignment_2/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsersSQL(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ageQuery := c.Query("age")
		sortBy := c.DefaultQuery("sort_by", "name")
		pageQuery := c.DefaultQuery("page", "1")
		pageSizeQuery := c.DefaultQuery("page_size", "10")

		page, err := strconv.Atoi(pageQuery)
		pageSize, err := strconv.Atoi(pageSizeQuery)
		offset := (page - 1) * pageSize

		query := "SELECT id, name, age FROM users WHERE 1=1"
		args := []interface{}{}

		if ageQuery != "" {
			query += " AND age = $1"
			age, _ := strconv.Atoi(ageQuery)
			args = append(args, age)
		}

		query += fmt.Sprintf(" ORDER BY %s LIMIT $2 OFFSET $3", sortBy)
		args = append(args, pageSize, offset)

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
			return
		}
		defer rows.Close()

		var users []model.User2
		for rows.Next() {
			var user model.User2
			if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
				return
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, users)
	}
}
