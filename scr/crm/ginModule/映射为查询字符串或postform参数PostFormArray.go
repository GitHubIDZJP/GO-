package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 映射为查询字符串或postform参数PostFormMap ---回传数组
func MapStrPostFormParamPostFormArray() {

	router := gin.Default()

	router.POST("/post", ginModel)
	router.Run(":8080")

}
func ginModel(c *gin.Context) {

	ids := c.PostFormArray("ids")
	prices := c.PostFormArray("prices")
	c.JSON(http.StatusOK, gin.H{
		"ids":    ids,
		"prices": prices,
	})

}

/*
接口:
http://localhost:8080/post?ids[0]=1&ids[1]=2
回传参数: 2次或者多次,回传值根据类型随便写
ids = 1
ids = 2
ids = 5

打印:
{
    "ids": [
        "1",
        "2",
        "5"
    ]
}

接口2: 加了个Prices：
http://localhost:8080/post?ids[0]=1&ids[1]=2&prices[0]=90&prices[1]=8
{
    "ids": [
        "10",
        "20"
    ],
    "prices": [
        "90",
        "8"
    ]
}
*/
