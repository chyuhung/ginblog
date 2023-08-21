package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length", "Authorization", "Content-Type"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return cors.New(config)
}
