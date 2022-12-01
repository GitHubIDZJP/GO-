package ginModule

import "github.com/gin-gonic/gin"

// 多部分/Urlen编码表单
func CodedForm() {
	//可以去"查询和发布表格"看下代码

	router := gin.Default()
	//c.PostForm(“message”) 获取表单中 message 的数据，如果没有message 则为空
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		//c.DefaultPostForm(“nick”, “anonymous”) 获取表单中 nick 的数据，如果没有设置一个默认值 anonymous；
		//测试
		nick := c.DefaultPostForm("sd", "") //c.DefaultPostForm("nick", "anonymous")
		/*
						   PostForm和DefaultPostForm的区别:
				一个有默认值，一个没有
			DefaultPostForm的默认值假设为空字符"",那postman里提交了也没任何效果，
		*/
		c.JSON(200, gin.H{

			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")

}

/*
post请求
接口:http://localhost:8080/form_post
{
    "message": "",
    "nick": "anonymous",
    "status": "posted"
}

*/
