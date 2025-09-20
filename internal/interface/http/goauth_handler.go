package http

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func OauthLogin(c *gin.Context) {
	provider := c.Param("provider") // "google" or "facebook"
	c.Request = c.Request.WithContext(
		context.WithValue(c.Request.Context(), "provider", provider),
	)
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func CallbackRoute(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(
		context.WithValue(c.Request.Context(), "provider", provider),
	)

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save user info in session
	session := sessions.Default(c)
	session.Set("user_name", user.Name)
	session.Set("user_email", user.Email)
	session.Save()

	// Redirect to profile or protected route
	c.Redirect(http.StatusFound, "/v1/api/user/")
}

func Logout(c *gin.Context) {
	_ = gothic.Logout(c.Writer, c.Request)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
