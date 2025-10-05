package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

	// Prepare JSON for POST to /v1/api/user/
	reqBody := map[string]interface{}{
		"username": user.Name,
		"name":     user.Name,
		"sex":      "unknown",
		"age":      0,
		"hashpass": "oauth_login",
		"email":    user.Email,
	}

	client := &http.Client{}

	// 1️⃣ Check if user already exists
	checkReq, err := http.NewRequest("GET", "http://localhost:8000/v1/api/user/email/"+user.Email, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create check request"})
		return
	}
	checkResp, err := client.Do(checkReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return
	}
	defer checkResp.Body.Close()

	if checkResp.StatusCode == http.StatusOK {
		// User already exists, skip creation
		c.JSON(http.StatusOK, gin.H{
			"message": "User already exists",
			"user":    reqBody,
			"status":  checkResp.Status,
		})
		return
	}

	// 2️⃣ User does not exist, create new one
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode JSON"})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/v1/api/user/", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create POST request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "API call failed", "body": string(body)})
		return
	}

	// Save user info in session
	session := sessions.Default(c)
	session.Set("user_name", user.Name)
	session.Set("user_email", user.Email)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User created via OAuth",
		"user":    reqBody,
		"status":  resp.Status,
	})
}

func GetUserDataRoute(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user_email")
	userName := session.Get("user_name")

	if userEmail == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userData := map[string]interface{}{
		"email": userEmail,
		"name":  userName,
	}
	fmt.Println(userData, "5555555555555555555555555555555")

	c.JSON(http.StatusOK, gin.H{"user": userData})
}

func Logout(c *gin.Context) {
	_ = gothic.Logout(c.Writer, c.Request)
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
