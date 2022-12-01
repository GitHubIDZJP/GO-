package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 从阅读器提供数据
func FeedDataFromTheReader() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.GET("/someDataFromReader", func(c *gin.Context) {
		/**
		https://raw.githubusercontent.com/gin-gonic/logo/master/color.png
		*/
		//用官网的urlStr01地址会报错503,用urlStr02可以正常获取到网址图片
		//var urlStr01 = "https://raw.githubusercontent.com/gin-gonic/logo/master/color.png"

		var urlStr02 = "https://golang.google.cn/images/gophers/ladder.svg"
		response, err := http.Get(urlStr02)
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}

/**
接口:
http://localhost:8080/someDataFromReader

*/
