package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "postgresql://postgres:cpre888@db.nreddhtmpvfcvmsktpai.supabase.co:5432/postgres"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	// // Optional: test the connection
	// if err := db.Ping(); err != nil {
	// 	log.Fatal("cannot ping database:", err)
	// }
	for i := 0; i < 10; i++ {
		if err := db.Ping(); err == nil {
			fmt.Println("Database connected")
			return db
		}
		log.Println("Database not ready, retrying...")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("cannot connect to database after retries")

	fmt.Println("Database connected")
	return db
}
