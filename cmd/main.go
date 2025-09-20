package main

import (
	"log"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/config"
	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	conn := db.Connect()
	defer conn.Close()

	config.Session_init()
	config.Goth_init()

	userRepo := repo.NewUserRepository(conn)
	userService := usecase.NewUserService(userRepo)

	http.SetUpRoutes(
		app,
		userService,
	)

	log.Fatal(app.Run(":8000"))
}
