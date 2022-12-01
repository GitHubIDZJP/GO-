package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
func AsciiJSON_Api() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"}) //设置代理，必须加上，不然下面会提示代理提示权限不安全
	r.GET("", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
		})

		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)

	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

//在终端输入http://localhost:8080/,显示404 page not found
/**

{
    "message": "请求成功"
}
{
    "lang": "GO语言",
    "tag": "<br>"
}

*/
