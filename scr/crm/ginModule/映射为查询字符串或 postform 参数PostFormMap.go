package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 映射为查询字符串或postform参数PostFormMap ---回传字典
func MapStrPostFormParam() {

	/**
	可以去"查询和发布表单"里看回传一个参数
	查询和发布表单:QueryAndPublishForms()

	*/
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		//ids_values := []int{1, 25, 7}
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids数组打印: %v; names数组打印: %v", ids, names)
		c.JSON(http.StatusOK, gin.H{

			"code": http.StatusOK,
			"ids":  ids,
			"name": names,
		})
		fmt.Println(names)
		c.String(200, "ok")

	})
	router.Run(":8080")
}

/***
接口:
http://localhost:8080/post?ids[a]=1234&ids[b]=hello
*/

/**
使用c.GetQueryMap("map")，则会返回map[1:v1 2:v2] true
/?name=Tom&name=Jack等于GetQueryArray

使用c.GetQueryMap("map")，则会返回map[1:v1 2:v2] true

如果使用的方法没有带Get前缀方法，那么返回对应的数据类型，如果对应的键不存在，那么则会返回一个空值，当然我们可以使用DefaultQuery指定缺省值
*/
