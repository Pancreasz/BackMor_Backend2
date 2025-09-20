package main

import (
	"log"
	"time"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/config"
	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️  No .env file found, relying on system env")
	}
	app := gin.Default()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your React app URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	conn := db.Connect()
	defer conn.Close()

	config.Session_init()
	config.Goth_init()

	userRepo := repo.NewUserRepository(conn)
	userService := usecase.NewUserService(userRepo)

	activityRepo := repo.NewActivityRepository(conn)
	activityService := usecase.NewActivityService(activityRepo)

	http.SetUpRoutes(
		app,
		userService,
		activityService,
	)

	log.Fatal(app.Run(":8000"))
}
