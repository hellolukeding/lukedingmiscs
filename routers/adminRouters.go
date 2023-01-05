package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//路由组
func AdminRoutersInit(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "admin user")
		})

		adminRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "admin news")
		})
	}
}
