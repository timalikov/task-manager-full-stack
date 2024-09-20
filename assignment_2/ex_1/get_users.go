package ex_1

import (
	"database/sql"
	"fmt"
	"log"
)

func GetUsers(db *sql.DB) {
	query := `
	SELECT * FROM users;`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}

		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error after reading rows: %v", err)
	}
}
