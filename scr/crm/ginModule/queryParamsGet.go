package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Gin中可以方便的获取URL中的查询参数，或者也可以简称为URL参数，
此类参数以?为起点，后面的k=v&k1=v1&k2=v2这样的字符串就是查询参数，
本小节通过示例演示获取URL查询参数的方式，以及背后的实现方式。
*/
func QueryParams() {
	var r = gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.GET("/", func(c *gin.Context) {
		var name = c.Query("name")
		var age = c.Query("age")
		//类似现在就已经带了参数，而不再需要客户端回传给自己
		//var name = c.DefaultQuery("name", "刘德华")
		//var age = c.DefaultQuery("age", "12")
		var data = map[string]interface{}{
			"name": name,
			"age":  age,
		}
		c.JSON(http.StatusOK, gin.H{

			"status":  200,
			"message": "响应成功",
			"data":    data,
		})
	})
	r.Run(":8080")
}

/*
*
回参前:http://localhost:8080

	{
	    "data": {
	        "age": "",
	        "name": ""
	    },
	    "message": "响应成功",
	    "status": 200
	}

回参后:http://localhost:8080?name=王强&age=20

	{
	    "data": {
	        "age": "20",
	        "name": "王强"
	    },
	    "message": "响应成功",
	    "status": 200
	}
*/
func DefaultQuery() {
	var r = gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.GET("/", func(c *gin.Context) {
		//var name = c.Query("name")
		//var age = c.Query("age")
		//类似现在就已经带了参数，而不再需要客户端回传给自己
		var name = c.DefaultQuery("name", "刘德华")
		var age = c.DefaultQuery("age", "12")
		var data = map[string]interface{}{
			"name": name,
			"age":  age,
		}
		c.JSON(http.StatusOK, gin.H{

			"status":  200,
			"message": "响应成功",
			"data":    data,
		})
	})
	r.Run(":8080")
}

/*
*
回传前: http://localhost:8080

	{
	    "data": {
	        "age": "12",
	        "name": "刘德华"
	    },
	    "message": "响应成功",
	    "status": 200
	}
*/
func queryGet() {
	var r = gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})

	r.GET("/", func(c *gin.Context) {

		name, ok := c.GetQuery("name")
		if !ok {
			name = "body"
		}
		age, ok := c.GetQuery("age")
		if !ok {
			age = "10"
		}

		//
		////var name = c.Query("name")
		////var age = c.Query("age")
		////类似现在就已经带了参数，而不再需要客户端回传给自己
		//var name = c.DefaultQuery("name", "刘德华")
		//var age = c.DefaultQuery("age", "12")
		//var data = map[string]interface{}{
		//	"name": name,
		//	"age":  age,
		//}
		c.JSON(http.StatusOK, gin.H{

			"status":  200,
			"message": "响应成功",
			"name":    name,
			"age":     age,
		})
	})
	r.Run(":8080")
}

/*
*
回传前: http://localhost:8080

	{
    "age": "10",
    "message": "响应成功",
    "name": "body",
    "status": 200
}
*/
