package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	"github.com/Pancreasz/BackMor_Backend2/infrastructure/router"
	http "github.com/Pancreasz/BackMor_Backend2/internal/interface/http"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"
)

func main() {
	conn := db.Connect()
	defer conn.Close()

	userRepo := repo.NewUserRepository(conn)
	userService := usecase.NewUserService(userRepo)

	activityRepo := repo.NewActivityRepository(conn)
	activityService := usecase.NewActivityService(activityRepo)

	// Create router with CORS + middleware
	app := router.NewRouter()

	http.SetUpRoutes(app, userService, activityService)

	// Print registered routes (for debugging)
	for _, r := range app.Routes() {
		fmt.Printf("%-6s %s\n", r.Method, r.Path)
	}

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(app.Run(":" + port))
}
