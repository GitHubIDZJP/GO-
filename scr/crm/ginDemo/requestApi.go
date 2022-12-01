package ginDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Api01() {
	//http://localhost:8080?name=sdsd
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.GET("/", func(context *gin.Context) {
		//要回传的参数
		name := context.Query("name")
		data := map[string]interface{}{
			"name": name,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "响应成功",
			"data":    data,
		})

	})
	r.Run(":8080")
}

/**
relativePath:只能能写/，不能写/ping,不然会报404
postman显示：
{
    "data": {
        "name": "sdsd"
    },
    "message": "响应成功",
    "status": 200
}


**/
