package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router.Use(gin.LoggerWithFormatter(logFormat))
func CustomHttpConfig() {

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusOK,
			"message": "大傻逼",
		})
	})
	//直接使用http.ListenAndServe() ------->自定义 HTTP 配置
	//http.ListenAndServe(":8080", router)
	router.Run(":8080")
}
