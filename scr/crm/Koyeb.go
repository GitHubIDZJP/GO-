package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

/*
*

使用 git 驱动部署在 Koyeb 无服务器平台上部署 Golang Gin 应用程序。

Gin是一个用 Golang 编写的网络框架，专注于性能和简单性。

通过使用 git 驱动部署在Koyeb上部署 Go Gin 应用程序，每次推送代码更改时，一旦构建和健康检查完成，您的应用程序将自动构建和提升。部署您的应用程序后，它将受益于 Koyeb 本机自动缩放、自动 HTTPS (SSL)、自动修复和零配置边缘网络的全局负载平衡
*/
func mainKoyeb() {
	//官网地址:--->
	//https://www.koyeb.com/tutorials/deploy-go-gin-on-koyeb
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
			"data":    "信息",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
