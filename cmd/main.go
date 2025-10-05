package main

import (
	"fmt"
	"log"

	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	conn := db.Connect()
	router := gin.Default()
	defer conn.Close()

	userRepo := repo.NewUserRepository(conn)
	userService := usecase.NewUserService(userRepo)

	activityRepo := repo.NewActivityRepository(conn)
	activityService := usecase.NewActivityService(activityRepo)

	http.SetUpRoutes(
		app,
		userService,
		activityService,
	)

	for _, r := range router.Routes() {
		fmt.Printf("%-6s %s\n", r.Method, r.Path)
	}
	log.Fatal(app.Run(":8000"))
}
