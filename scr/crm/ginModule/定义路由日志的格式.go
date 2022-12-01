package ginModule

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
路由的默认日志是
POST   /foo                      --> main.main.func1 (3 handlers)
GET    /bar                      --> main.main.func2 (3 handlers)
GET    /status                   --> main.main.func3 (3 handlers)
*/

// 定义路由日志的格式
func DefineFormatofTheRouteLog() {
	r := gin.Default()
	/*

		给定格式（例如 JSON、键值或其他格式）记录此信息，则可以使用gin.DebugPrintRouteFunc



	*/

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})
	/*
			接口:http://localhost:8080/foo
		打印:foo

	*/

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})
	/*
			接口:http://localhost:8080/bar
		    打印:bar

	*/

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	/*
			接口:http://localhost:8080/status
		    打印:"ok"

	*/
	r.Run(":8080")
}

/*
终端打印:

2022/11/29 02:52:03 endpoint POST /foo crm/ginModule.DefineFormatofTheRouteLog.func2 3
2022/11/29 02:52:03 endpoint GET /bar crm/ginModule.DefineFormatofTheRouteLog.func3 3
2022/11/29 02:52:03 endpoint GET /status crm/ginModule.DefineFormatofTheRouteLog.func4 3


*/
