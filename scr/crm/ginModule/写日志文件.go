package ginModule

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// 写日志文件
func WritingLog() {
	//禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	//登录到文件
	f, _ := os.Create("gin.log")
	//如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")

}

/*
get接口:
http://localhost:8080/ping

打印:

pong


*/
