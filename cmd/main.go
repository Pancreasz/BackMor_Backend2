package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"

	"github.com/gofiber/fiber/v2"
)

func main() {

	conn := db.Connect()
	defer conn.Close()

	repo := repo.NewUserRepository(conn)
	ctx := context.Background()

	user, err := repo.GetByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user: %+v\n", user)

	app := fiber.New()
	http.SetUpCustomerRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
