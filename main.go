package main

import (
	"leargin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//设置路由
	r := gin.Default()

	//信任代理
	r.SetTrustedProxies([]string{"192.168.0.102", "192.168.0.19"})

	//router group
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.ApiTestInit(r)

	r.Run(":8080")
}
