package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return cors.Default()
	// return func(c *gin.Context) {
	// 	cors.New(cors.Config{
	// 		//AllowAllOrigins:  true,
	// 		AllowOrigins:     []string{"*"},
	// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 		AllowHeaders:     []string{"*"},
	// 		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
	// 		AllowCredentials: true,
	// 		MaxAge:           12 * time.Hour,
	// 	})
	// }
}
