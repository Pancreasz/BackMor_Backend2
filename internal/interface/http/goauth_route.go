package http

import (
	"github.com/gin-gonic/gin"
)

func SetUpOAuthRoutes(api gin.IRouter) {
	OAuthRoutes := api.Group("/auth")
	OAuthRoutes.GET("/:provider", OauthLogin)
	OAuthRoutes.GET("/:provider/callback", CallbackRoute)
	OAuthRoutes.GET("/auth/callback", GetUserDataRoute)
	OAuthRoutes.GET("/logout", Logout)
}
