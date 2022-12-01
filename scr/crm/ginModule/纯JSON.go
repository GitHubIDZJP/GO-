package ginModule

import (
	"github.com/gin-gonic/gin"
)

/*
通常，JSON 用它们的 unicode 实体替换特殊的 HTML 字符，例如<变成\u003c. 如果您想按字面意思对此类字符进行编码，则可以改用 PureJSON。此功能在 Go 1.6 及更低版本中不可用

*/
// 纯JSON  在postman请求解析是显示一样的，但是直接把api复制到谷歌浏览器地址就能显示不一样的body了
func PureJSON() {
	//提供 unicode 实体
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	/*
				Api:
				http://localhost:8080/json
			body:
			{
			    "html": "\u003cb\u003e Hello, world!\u003c/b\u003e"
			}
		普通json 会将< > 特殊字符转为 unicode
	*/

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	r.Run(":8080")
}

/**
Api:
		http://localhost:8080/json
	body:
	{
	    "html": "<b>Hello, world!</b>"
	}
    pureJson 将会保留特殊字符
*/
