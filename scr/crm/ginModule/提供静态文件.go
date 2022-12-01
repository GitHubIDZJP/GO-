package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 加载模版里的HTML文件
func ProvideStaticFile() {

	/**
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.Run(":8080")
	*/

	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	//router.LoadHTMLFiles("templates/") ////加载templates目录下所有文件
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "我是大傻逼",
		})
	})
	router.Run(":8080")
	/**
	接口:http://localhost:8080/index
	放到postman请求呢 ，则显示整个HTML网页内容
	放到浏览器则显示body里的元素内容



	*/
}
