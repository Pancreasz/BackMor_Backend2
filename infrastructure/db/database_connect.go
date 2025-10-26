package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// Read password from environment variable
	dbPassword := os.Getenv("PAINAI_DB_PASSWORD")
	// fmt.Println(dbPassword, 12312312312312312)
	if dbPassword == "" {
		log.Fatal("PAINAI_DB_PASSWORD environment variable is not set")
	}

	connStr := fmt.Sprintf(
		"postgresql://postgres:%s@backmor-database.postgres.database.azure.com:5432/postgres?sslmode=require",
		dbPassword,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	for i := 0; i < 10; i++ {
		if err := db.Ping(); err == nil {
			fmt.Println("Database connected")
			return db
		}
		log.Println("Database not ready, retrying...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("cannot connect to database after retries")
	return db
}
