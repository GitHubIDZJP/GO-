package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/fvbock/endless"

func EndlessShutDownService() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.SecureJSON(http.StatusOK, gin.H{
			"code":     http.StatusOK,
			"服务请求重新开启": "将要执行服务开启",
		})
	})
	//c.String(200, "test shutdown!")
	err := endless.ListenAndServe("localhost:8080", r)
	print(err)

}
