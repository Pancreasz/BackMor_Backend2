package main

import (
	"log"

	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conn := db.Connect()
	defer conn.Close()

	userRepo := repo.NewUserRepository(conn)
	userService := usecase.NewUserService(userRepo)

	http.SetUpRoutes(
		app,
		userService,
	)

	log.Fatal(app.Listen(":8000"))
}
