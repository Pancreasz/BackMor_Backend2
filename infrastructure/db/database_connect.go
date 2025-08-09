package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "postgresql://postgres:cpre888@localhost:5432/backmor_database?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	// Optional: test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("cannot ping database:", err)
	}

	fmt.Println("Database connected")
	return db
}
