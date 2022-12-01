package ginDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 基础请求框架
func BashRequestFramework() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"}) //设置代理，必须加上，不然下面会提示代理提示权限不安全
	r.GET("/", func(obj *gin.Context) {
		//删除json结果给调用方
		obj.JSON(http.StatusOK, gin.H{
			"messgae": "请求成功",
		})
	})
	r.Run(":8080")
}
