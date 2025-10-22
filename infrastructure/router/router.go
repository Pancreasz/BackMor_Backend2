package router

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Read allowed origins from env
	allowed := os.Getenv("ALLOWED_ORIGINS")
	var allowOrigins []string
	if allowed == "" {
		allowOrigins = []string{"http://localhost:3000"} // dev default
	} else {
		for _, o := range strings.Split(allowed, ",") {
			if s := strings.TrimSpace(o); s != "" {
				allowOrigins = append(allowOrigins, s)
			}
		}
	}

	cfg := cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Location", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(cfg))

	return r
}
