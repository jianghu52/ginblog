package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	router := r.Group("api/v1")
	router.GET("hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	r.Run(utils.HttpPort)
}
