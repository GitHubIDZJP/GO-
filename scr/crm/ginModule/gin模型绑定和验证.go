package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 回传的参数用户名和密码必须以小字母开头(user和password)
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// 模型绑定和验证
func Model_binding_and_validation() {
	//如果存在绑定错误，请求将被中止
	router := gin.Default()

	//1  绑定示例 JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "未授权"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "您已登录"})

		/*
			接口:http://localhost:8080/loginJSON?user=manu&password=123

						{
						    "error": "EOF"
						}

		*/
	})

	//2
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			/*
					绑定xml:-->loginXML
				http://localhost:8080/loginXML?user=manu&password=123
					{
					    "error": "EOF"
					}

			*/
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "未授权"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "您已登录"})
	})

	//3 绑定 HTML 表单的示例 (user=manu&password=123)--->这有这个方法是行得通的
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 这将根据内容类型标头推断要使用的绑定程序。
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "未经授权"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "您已登录了呢"})
		/*
							接口:
				//密码和用户名对时
			密码: 123
					http://localhost:8080/loginForm?user=manu&password=123
					{
					    "status": "您已登录了呢"
					}
				//密码和用户名错时:
			密码:124
				http://localhost:8080/loginForm?user=manu&password=124
			{
			    "status": "未经授权"
			}

		*/
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
