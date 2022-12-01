package ginModule

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// XML/JSON/YAML/ProtoBuf 渲染
func RenderingMain() {
	//r := gin.Default()
	//，
	r := gin.New()
	/**
	默认没有中间件
	用
	r := gin.New()
	替换：
	r := gin.Default()
	*/

	// gin.H 是一个捷径 map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "hey",
		})
	})
	/*
		{
		    "message": "hey",
		    "status": 200
		}
	*/

	r.GET("/moreJSON", func(c *gin.Context) {
		// 还可以使用结构
		var msg struct {
			Name    string
			Message string
			Number  int
		}

		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意msg。Name在JSON中变成了“user”
		// 将会输出  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})
	/**
	接口：http://localhost:8080/moreJSON
	json形式输出
	{
	    "user": "Lena",
	    "Message": "hey",
	    "Number": 123
	}

	*/

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})
	/**
	接口：http://localhost:8080/someXML
	xml形式输出
	<map>
	    <message>hey</message>
	    <status>200</status>
	</map>

	*/

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})
	/**
	接口：http://localhost:8080/someYAM
	YAML形式输出
	message: hey
	status: 200


	*/

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf的具体定义写在testdata/protoexample文件中.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 注意，数据在响应中变成了二进制数据
		//将输出protoexample。测试protobuf序列化数据
		c.ProtoBuf(http.StatusOK, data)
	})
	/**
	接口：http://localhost:8080/someProtoBuf
	someProtoBuf形式输出
	test


	*/

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
