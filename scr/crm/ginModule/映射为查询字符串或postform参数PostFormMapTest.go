package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type m_u struct {
	Age  string
	Name string
}

// 映射为查询字符串或postform参数PostFormMap ---回传字典
func MapStrPostFormParamTest() {

	/**
	可以去"查询和发布表单"里看回传一个参数
	查询和发布表单:QueryAndPublishForms()

	*/
	router := gin.Default()

	router.POST("/users", PostFormMap)
	router.Run(":8080")

}

// 回调
func PostFormMap(c *gin.Context) {

	user := c.PostFormMap("user")
	//var s = m_u{Age: user}
	ageT := c.Query("age")
	name := c.Query("name")
	m_dataInit := m_u{ageT, name}
	c.JSON(http.StatusOK, gin.H{
		"users": user,
		"model": m_dataInit,
		//"userDic": m_dataInit,
	})
}

//打印{"user":{"age":"20","name":"juanmaofeifei"}}
/*
http://localhost:8080/post

?user[id]=111&user[name]=张三
回参数组
http://localhost:8080/post?user[id]=111&user[name]=张三
{
    "user": {}  这里是PostFormMap---map,所以打印是{},如果是PostFormArray,则打印[]
}
*/
