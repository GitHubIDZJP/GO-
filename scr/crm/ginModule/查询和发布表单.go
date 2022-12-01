package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
Query 和 post form 翻译后查询和发布表单


*/

/*
可以查看go函数
1.postRequest()
2.queryParams()
查询和发布表单

	query 查询
	DefaultQuery
	PostForm发布表单
*/
func QueryAndPublishForms() {
	router := gin.Default()
	//GET请求 url ?后面是querystring参数,key = value格式,多个key-value用&连接
	router.POST("/post", func(c *gin.Context) {
		///queryParams?name=randySun&age=18
		id := c.Query("id") // 查询参数：放在url里的用query
		////获取表单中 "page 的数据，如果没有设置一个默认值 0；
		page := c.DefaultQuery("page", "0")
		// c.PostForm("name")获取表单中 name 的数据，如果没有name 则为空
		name := c.PostForm("name") //重点: body 里的x-www-form-urlencoded 参数回参用PostForm

		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.SecureJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "回传成功",
			//这下面是为了方便查看传回来的参数值，实际开发了前端传回来后直接显示回传成功code：200就行
			"id":   id,
			"page": page,
			"name": name,
		})
	})
	router.Run(":8080")

}

/*
post接口:
http://localhost:8080/post?id=1234&page=0&name=TestName&message=meg

回传成功的参数:
在终端可以看到
/post?id=1234&page=0&name=TestName&message=msg"

接口:
http://localhost:8080/post
{
    "code": 200,
    "id": "",
    "message": "回传成功",
    "name": "",
    "page": "0"
}

postman里回参步骤:
1. 在Params回参id和phge分别对应上面的c.Query("id")
2.在body下x-www-form-urlencoded回参name和message 分别对应上面的c.PostForm("message")



*/

/**


如下面官方的例子中 ?id=1234&page=1 这个，就是 查询参数(query params)
你会看到这个参数就放在url里的

如果参数不是放在url里的，也可以在body里，比如 body 里的x-www-form-urlencoded 参数，如下面的name=manu&message=this_is_great
对于gin，要使用 name := c.PostForm("name") api

关注下 Content-Type 这个字段，表示了body的类型，更多看看这篇文章
https://imququ.com/post/four-...

还有像 c.FormFile，用于处理上传文件的

*/
