package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	router.GET("hello", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})
	r.Run(utils.HttpPort)
}
