package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func postRequest() {
	type Info struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	//post请求

	r.POST("/", func(c *gin.Context) {

		// 以Info为模板初始化data
		var data Info

		// 将请求参数绑定到data
		c.BindJSON(&data)

		fmt.Println("=data=>>", data)

		c.JSON(200, gin.H{
			"status":  1000,
			"message": "返回电影名称和评分",
			"data":    data,
		})

	})

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
	//r.Run()
}

/*
*
POST回传前: http://localhost:8080/movie

	{
    "data": {
        "name": "",
        "score": 0
    },
    "message": "响应成功",
    "status": 200
}
POST回传后: http://localhost:8080

	{
    "age": "10",
    "message": "响应成功",
    "name": "body",
    "status": 200
}
*/
