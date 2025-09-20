package http

import (
	"github.com/Pancreasz/BackMor_Backend2/infrastructure/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func SetUpRoutes(
	app *gin.Engine,
	userService UserService,
) {

	store := cookie.NewStore(config.HashKey, config.BlockKey)
	app.Use(sessions.Sessions("mysession", store))
	gothic.Store = store
	api := app.Group("/v1/api")
	SetUpUserRoutes(api, userService)
	SetUpOAuthRoutes(api)
}
