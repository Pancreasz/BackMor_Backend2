package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	HashKey  []byte
	BlockKey []byte
)

func Session_init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get session keys from env
	HashKey = []byte(os.Getenv("SESSION_HASH_KEY"))
	BlockKey = []byte(os.Getenv("SESSION_BLOCK_KEY"))

	if len(HashKey) == 0 {
		log.Fatal("SESSION_HASH_KEY must be set in environment variables")
	}
	log.Printf("HashKey: %d bytes, BlockKey: %d bytes\n", HashKey, BlockKey)

}
