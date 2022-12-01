package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func pingResponse(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 自定义日志文件函数
func logFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,                       //客户端ip
		param.TimeStamp.Format(time.RFC1123), //时间戳--国际标准 RFC3339
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,             //等待时间
		param.Request.UserAgent(), //用户代理
		param.ErrorMessage,
	)
}

// 自定义日志文件
func CustomLogFile() {
	router := gin.New()
	//LoggerWithFormatter中间件将把日志写入gin.DefaultWriter
	// 默认为 gin.DefaultWriter = os.Stdout
	//单独把方法抽出来用func函数调用
	/*	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// 自定义格式
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,                       //客户端ip
				param.TimeStamp.Format(time.RFC1123), //时间戳
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,             //等待时间
				param.Request.UserAgent(), //用户代理
				param.ErrorMessage,
			)
		}))
	*/
	//调用自定义日志文件
	router.Use(gin.LoggerWithFormatter(logFormat))

	router.Use(gin.Recovery())
	/*router.GET("/", func(c *gin.Context) {
		//c.String(200, "pong")
		c.SecureJSON(http.StatusOK, gin.H{
			"1": "哈哈哈",
		})
	})

	*/

	router.GET("/ping", pingResponse)
	router.Run(":8080")
}

/*
https://blog.csdn.net/wohu1104/article/details/126689036
百度的goland打印:
127.0.0.1 - [2022-07-13T17:16:31.300579298+08:00] "GET /ping HTTP/1.1 200 20.821µs "curl/7.58.0" "
自己Mac下goland终端打印:
[Tue, 29 Nov 2022 16:49:27 CST] "GET /ping HTTP/1.1 200 20.129µs "PostmanRuntime/7.29.2" "

解析:
param.ClientIP, -------------->      127.0.0.1
		param.TimeStamp.Format(time.RFC1123), --------------> /2022-07-13T17:16:31.300579298+08:00
		param.Method, --------------> GET
		param.Path, --------------> /ping
		param.Request.Proto, --------------> HTTP/1.1
		param.StatusCode, --------------> 200
		param.Latency,    --------------> 20.821µs         //等待时间
		param.Request.UserAgent(), --------------> //用户代理 curl/7.58.0
		param.ErrorMessage, --------------> code为200时没任何显示


router.Use(gin.LoggerWithFormatter(logFormat))这方法不调用时，则postman测试请求接口时则啥也不打印

*/
