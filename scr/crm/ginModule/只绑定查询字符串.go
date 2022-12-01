package ginModule

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func BindOnlyInQueryString() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	//1ShouldBindQuery函数只绑定查询参数而不绑定发布数据
	//bindQuery 也适用于“form”标签

	/*
		都可以查询绑定的参数字符串
		if c.ShouldBindQuery(&person) == nil {}
		if c.BindQuery(&person) == nil {}
		if c.Bind(&person) == nil {}
	*/

	//if c.BindQuery(&person) == nil {
	//if c.Bind(&person) == nil {
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== 仅按查询字符串绑定 ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "操作成功了呢ss")
}

/*
查询只绑定的参数字符串：
name和address
http://localhost:8080/testing?name=eason&address=xyz
操作成功了呢
是因为里面参数里包括了name和address
*/

// 尝试c.BindJSONJSON 数据
type PersonTest struct {
	//把原来的form换为josn
	Name    string `json:"name"`
	Address string `json:"address"`
}

func BindOnlyInQueryStringTest() {
	route := gin.Default()
	route.Any("/testing", startPageTest)
	route.Run(":8080")
}
func startPageTest(c *gin.Context) {
	var p PersonTest
	/**
	if c.ShouldBindQuery(&p) == nil { }报错
	if c.Bind(&p) == nil {} 不报错
	if c.BindQuery(&p) == nil {} 不报错
	*/

	if c.BindQuery(&p) == nil {
		log.Println("====== 仅按查询字符串绑定 ======")
		log.Println(p.Name)
		log.Println(p.Address)
	}
	c.String(200, "操作成功了呢sssdsd")
}
