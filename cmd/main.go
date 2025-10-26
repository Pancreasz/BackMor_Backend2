package main

import (
	"log"

	"github.com/Pancreasz/BackMor_Backend2/infrastructure/config"
	"github.com/Pancreasz/BackMor_Backend2/infrastructure/db"
	"github.com/Pancreasz/BackMor_Backend2/infrastructure/router"
	"github.com/Pancreasz/BackMor_Backend2/internal/interface/http"
	repo "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {

	app := router.NewRouter()

	conn := db.Connect()
	defer conn.Close()

	// config.InitFirebase()
	config.Session_init()
	config.Goth_init()

	store := cookie.NewStore(config.HashKey, config.BlockKey)
	app.Use(sessions.Sessions("mysession", store))

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
