package ex_3

import (
	"assignment_2/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsersDirect(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []model.User
		for rows.Next() {
			var user model.User
			err := rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}

		c.JSON(http.StatusOK, users)
	}
}

func CreateUserDirect(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
		err := db.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, user)
	}
}

func UpdateUserDirect(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "UPDATE users SET name=$1, age=$2 WHERE id=$3"
		_, err := db.Exec(query, user.Name, user.Age, id)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, user)
	}
}

func DeleteUserDirect(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		query := "DELETE FROM users WHERE id=$1"
		_, err := db.Exec(query, id)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %d deleted", id)})
	}
}
