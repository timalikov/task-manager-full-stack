package ex_1

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Table created successfully")
}
