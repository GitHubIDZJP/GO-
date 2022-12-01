package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HTML 渲染
func HTMLRendering() {
	//可以先看"加载模版里的HTML文件"
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	//router.LoadHTMLFiles("templates/") ////加载templates目录下所有文件
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "我是大傻逼",
		})
	})
	router.Run(":8080")
}
