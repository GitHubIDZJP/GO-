package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 分组路线
func FroupRouter() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint1)
		v1.POST("/submit", submitEndpoint1)
		v1.POST("/read", readEndpoint1)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint2)
		v2.POST("/submit", submitEndpoint2)
		v2.POST("/read", readEndpoint2)
	}

	router.Run(":8080")
}

func loginEndpoint1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"login": "登录1",
	})
}
func loginEndpoint2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"login": "登录2",
	})
}

// 提交断点
func submitEndpoint1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"submit": "提交数据",
	})
}
func submitEndpoint2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"submit": "提交数据2",
	})
}

// 读取端点
func readEndpoint1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"readEndpoint": "请求数据点",
	})

}
func readEndpoint2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"readEndpoint": "请求数据点",
	})

}

/*

router := gin.New()

router.GET("/testGet",func(c *gin.Context){
    //处理逻辑
})

router.POST("/testPost",func(c *gin.Context){
    //处理逻辑
})

router.PUT("/testPut",func(c *gin.Context){
    //处理逻辑
})

router.DELETE("/testDelete",func(c *gin.Context){
    //处理逻辑
})

router.PATCH("/testPatch",func(c *gin.Context){
    //处理逻辑
})

router.OPTIONS("/testOptions",func(c *gin.Context){
    //处理逻辑
})

router.OPTIONS("/testHead",func(c *gin.Context){
    //处理逻辑
})




*/
