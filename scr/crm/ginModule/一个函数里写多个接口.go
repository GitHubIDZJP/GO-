package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
)

// 一个接口里多个Api
func MultipleApi() {

	//接口1
	r := gin.Default()

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"Api":  "接口000000000001",
		})
	})

	//接口2

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": http.StatusOK,
			"Api":  "接口00000000002",
		})
	})

	//接口3
	r.POST("/returnParam", returnUserName)
	/*
		postman请求：
			get：正常显示
			post：正常显示
		谷歌浏览器放地址栏请求:
			get:404页面不存在
		    post：正常显示
	*/

	//接口4 和4是一对，一个接收
	r.GET("/user", InitStruct_json)
	//接口5 和4是一对，一个给客户端请求
	r.GET("/get", getU)

	r.Run(":8080")
	//ginModule.RenderingMain()
}

type Pe struct {
	Name string `form:"name"`
}

// 3
func returnUserName(c *gin.Context) {

	//func(c *gin.Context) {

	//var normal string = "回传成功"
	//需要接收的参数名
	userName := c.Query("username")
	//userName := c.PostForm("username")
	/*
		c.Query：
		是在postman下Params填写回传参数
		c.PostForm
		是在postman下X-www-Form-urlenced填写回传参数

	*/

	//userName := c.QueryMap()
	data := Pe{
		Name: userName,
		//"撒大声地",
	}
	fmt.Println("打印纸:\n", data.Name)
	var fetchName string = data.Name

	res, err := ffjson.Marshal(data)
	if err != nil {
		fmt.Println("格式化错误")
		fmt.Println(err.Error())
		return
	}
	// 得到是字节数组，所以还有转为string
	fmt.Println("sdssdd \n", string(res))

	//userName := c.PostForm("username")
	//userName := c.Param("username")//json里显示不出来-少用,因为这是错误的

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"name": fetchName, //这个参数不需要的，只是为了方便查看而已

	})

}

type User_TTTT struct {
	Username string ``
	Password string ``
	Email    string
}

var me User_TTTT

// 4
// 获取前端和客户端回传的参数值
func InitStruct_json(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	me = User_TTTT{username,
		password,
		email,
	}
	err := c.ShouldBind(&me)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": me, //这里只是为了方便演示看是否真的传进来了，真正开发里值直接判断code是否为200和直接返回回传成功
		})
	}
}

// 把从前端或客户端获取的值以json形式返还给客户端
func getU(c *gin.Context) {
	name, ok := c.GetQuery("username")
	if !ok {
		name = me.Username
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"name":     name,
		"password": me.Password,
		"age":      me.Email,
	})

}
