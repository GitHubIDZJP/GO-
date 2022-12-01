package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONP() {

	//使用 JSONP 从不同域中(跨域)的服务器请求数据。如果查询参数回调存在，则将回调添加到响应主体

	r := gin.Default()
	r.GET("/JSONP", CallBack)

	r.Run(":8080")

}

func CallBack(c *gin.Context) {
	/*

		interface{} 可以代表任意类型
		interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型

	*/
	type Recep struct {
		Name, Gift string
		Attended   bool
	}

	//数组结构体
	var reslice = []map[string]interface{}{
		{"Name": "张三", "Gift": "李四", "Attended": false},
		{"Name": "哈哈", "Gift": "微微", "Attended": true},
	}

	//c.JSON(http.StatusOK, data) //不支持跨域
	c.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": reslice,
	}) //支持跨域，最好用这个
}

/**
get请求
接口:http://localhost:8080/JSONP
打印:
{
    "code": 200,
    "data": [
        {
            "Attended": false,
            "Gift": "李四",
            "Name": "张三"
        },
        {
            "Attended": true,
            "Gift": "微微",
            "Name": "哈哈"
        }
    ]
}
*/

/**
json数据格式在前端开发这个领域用的比较多，其优点在于
 （1）基于纯文本，跨平台传递简单；

（2）JavaScript原生支持，后台语言几乎全部支持；

（3）轻量级数据格式，占用字符数量极少，特别适合互联网传递；

（4）可读性较强；

（5）容易编写和解析；

 (6) 不支持跨域

 (7) 一种数据交换格式


jsonp是一个跨域交互协议，可以理解为jsonp约定了json的这个数据怎样进行传递。简单的说，
 （1）支持跨域教育

区别:
1 JSON是一种数据交换格式，而JSONP是一种依靠开发人员的聪明才智创造出的一种非官方跨域数据交互协议。一个是描述信息的格式，一个是信息传递双方约定的方法。

2 一句话就是说：json返回的是一串数据；而jsonp返回的是脚本代码（包含一个函数调用
*/

/*

{
    "code": 200,
    "data": {
        "arr": [
            10,
            20,
            30,
            90
        ],
        "foo": "bar"
    }
}

*/
