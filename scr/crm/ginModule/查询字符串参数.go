package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询字符串参数
func SearchStrParameter() {
	router := gin.Default()

	// 使用现有的底层请求对象解析查询字符串参数.
	// 请求响应匹配的url:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 快捷方式 c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "查询 %s %s", firstname, lastname)
	})
	router.Run(":8080")
}

/***

http://localhost:8080/welcome

http://localhost:8080/welcome?firstname=Jane&lastname=Doe

*/
