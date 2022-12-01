package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecureJSON() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// r.SecureJsonPrefix(")]}',\n")

	//使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
	r.GET("/someJSON01", func(c *gin.Context) {
		//数组
		//vary 数组名 [数组元素长度] 数组类型
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
		//c.JSON(http.StatusOK, names)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

/*
c.SecureJSON(http.StatusOK, names)
打印:
while(1);[
    "lena",
    "austin",
    "foo"
]

c.JSON(http.StatusOK, names)
打印:
[
    "lena",
    "austin",
    "foo"
]
*/
