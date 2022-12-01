package ginModule

import "github.com/gin-gonic/gin"

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// gin框架中的 basicauth 认证模块，通过BasicAuth中间件可以快速实现 http 基础认证，提高应用安全性
func CEMTERComMain() {
	r := gin.Default()
	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(200, gin.H{
				"user":   user,
				"secret": secret,
			})
		} else {
			c.JSON(200, gin.H{
				"user":   user,
				"secret": "No Secret",
			})
		}
	})
	r.Run(":8080")
}

/**
http://127.0.0.1:8080/admin/secrets
# 输出入账号/密码 foo/bar
{
    "secret": {
        "email": "foo@bar.com",
        "phone": "123433"
    },
    "user": "foo"
}



*/
