package ginModule

import "github.com/gin-gonic/gin"

type LoginFormModel struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 多部分/URL 编码绑定
func URLCodingBind() {

	//可以去"只绑定查询字符串"看下代码

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 可以使用显式绑定声明绑定多部分表单:
		// c.ShouldBindWith(&form, binding.Form)
		// 或者您可以简单地使用自动绑定 ShouldBind method:
		var form LoginFormModel
		// 在这种情况下，将自动选择适当的绑定
		if c.ShouldBind(&form) == nil {
			//if form.User != "manu" || form.Password != "123" {
			if form.User == "zjp" && form.Password == "123456" {
				var userName = "您的用户名是:" + form.User
				c.JSON(200, gin.H{"status": userName})
			} else {
				c.JSON(401, gin.H{"status": "未授权"})
			}
		}
	})
	router.Run(":8080")

}

/**
接口:
code:401
http://localhost:8080/login?user=12&password=23(密码或者用户名错误时)
{
    "status": "未授权"
}
code:200
用户名: zjp
密码: 123456
http://localhost:8080/login?user=zjp&password=123456
{
    "status": "您的用户名是:zjp"
}
*/
