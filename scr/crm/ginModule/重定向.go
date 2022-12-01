package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 1 重定向
func Redirection() {
	e := gin.New()
	//e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		//指定跳转到某个网页
		//googleUrlStr = "http://www.google.com/"
		baiduUrlStr := "http://wwww.baidu.com"
		//发出 HTTP 重定向很容易。支持内部和外部位置。
		c.Redirect(http.StatusMovedPermanently, baiduUrlStr)

	})
	e.Run(":8080")

}

// 2 从 POST 发出 HTTP 重定向
func PostSendHttpRedirection() {
	e := gin.New()
	e.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "http://wwww.baidu.com")
	})
	e.Run(":8080")
}

//3 发出路由器重定向，使用HandleContext

func SenderRouterRedirection() {
	r := gin.New()
	r.GET("/test", func(c *gin.Context) {

		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	r.Run(":8080")
	/**
	接口参数test还是test2,body结果都是一样，通过HandleContext多层绑定路由指向
	{
	    "hello": "world"
	}


	*/
}
