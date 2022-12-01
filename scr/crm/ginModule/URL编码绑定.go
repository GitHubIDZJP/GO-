package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	//其中 form：“user” 表示前端提交form表单时User对应的key的名称为：user
	User string `form:"user" binding:"required" json:"user"`
	//其中 form：“password” 表示前端提交form表单时Password对应的key的名称为：password
	Password string `form:"password" binding:"required" json:"password"`
}

// 1未绑定
func MultipartUrl_Not_Bind() {

	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"name ": name,
			"age":   age,
		})
	})

	//r.Run(":8080")
	/*
				接口：
				http://localhost:8080/user/:name/:age或者http://localhost:8080/user/:name/:age?name=jk&age=34
			打印:
		{
		    "age": ":age",
		    "name ": ":name"
		}
	*/

}

// 2 Multipart/Urlencoded 绑定---ShouldBind
func MultipartUrlBind() {

	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 可以使用显式绑定声明绑定多部分表单:
		//	 c.ShouldBindWith(&form, binding.Form)
		//或者您可以简单地使用ShouldBind方法自动绑定:
		var form LoginForm
		// 在这种情况下，将自动选择适当的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "zjp" && form.Password == "123" {
				c.JSON(200, gin.H{
					"status": "登录了",
					"用户":     form.User,
					"密码":     form.Password,
				})
			} else {
				c.JSON(401, gin.H{"status": "未经许可"})
			}
		}
	})
	router.Run(":8080")
	/*
			http://localhost:8080/login?user=zjp&password=123
					打印:
					{
					    "status": "登录了"
					}
				http://localhost:8080/login?user=1&password=23
		打印:
			{
			    "status": "未经许可"
			}
	*/
}

// 3
//type Login_B struct {
//	User     string `form:"user" json:"user"`
//	Password string `form:"password" json:"password"`
//}
