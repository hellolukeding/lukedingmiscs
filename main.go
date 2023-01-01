package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//设置路由
	r := gin.Default()

	//信任代理
	r.SetTrustedProxies([]string{"192.168.0.102"})

	// r.GET("/", func(c *gin.Context) {
	// 	// TODO
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world",
	// 	})
	// })

	//路由传参
	r.GET("/news", func(c *gin.Context) {
		aid := c.Query("aid")
		c.String(200, "aid=%s", aid)
	})

	//动态路由
	r.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.String(200, "userID=%s", uid)
	})

	//c.String() c.JSON() c.JSONP() c.XML() c.HTML()

	//返回一个JSON
	r.GET("/someJSON", func(c *gin.Context) {
		//自行拼接json
		// gin.H 是 map[string]interface{}的缩写
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})

	//返回结构体构建的JSON
	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}

		msg.Name = "name"
		msg.Message = "success"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})

	//JSONP
	r.GET("/someJSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	//返回一个XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})

	r.GET("/moreXML", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "test"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})

	//返回一个HTML
	// 	Gin 框架中使用 c.HTML 可以渲染模板，渲染模板前需要使用 LoadHTMLGlob()或者
	// LoadHTMLFiles()方法加载模板。
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
	})
	r.Run(":8080")
}
