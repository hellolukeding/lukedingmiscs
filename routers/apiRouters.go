package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"username": "张三",
				"age":      18,
			})
		})

		apiRouter.GET("/news", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"title": "这是新闻"})
		})
	}
}
