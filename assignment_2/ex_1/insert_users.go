package ex_1

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertUser(db *sql.DB, name string, age int) {
	query := `
	INSERT INTO users (name, age)
	VALUES ($1, $2)
	RETURNING id;`

	var id int
	err := db.QueryRow(query, name, age).Scan(&id)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	fmt.Printf("Inserted user with ID: %d\n", id)
}
