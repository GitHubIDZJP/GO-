package ginModule

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//中间件内的协程
/*

在一个中间件 或 处理器 中 启动一个新的协成时，不能使用它里面原始

 的 context ， 只能使用它的只读副本  cc = c.Copy

gin框架的每一个request请求都会开一个goroutine来处理，所以天然支持高并发
*/
func CoroutinesWithinMiddleware() {
	r := gin.Default()

	// 协程-异步(async) 测试
	r.GET("/long_async", func(c *gin.Context) {

		// 创建要在goroutine中使用的副
		cp := c.Copy()
		// golang通过goroutine实现高并发,而开启goroutine的钥匙正是go关键字
		//go func(){}()  以协程方式运行,这样就可以提供程序的并发处理能力
		go func() {
			//用时间模拟一个长任务 延迟5秒后执行
			time.Sleep(5 * time.Second) //定时器
			//这里使用你创建的副本：注意这里使用的是上下文对象的只读副本 "cCp"
			log.Println("完成了!在路径" + cp.Request.URL.Path)
		}()

		c.SecureJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "网络请求",
			"请求路径":    c.Request.URL.Path,
			"协程":      "异步步",
		})
	})
	//携程-->异步(async) 测试
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second) //定时器
		//这里没有使用协程，所以可以直接使用上下文对象 c
		log.Println("完成了!在路径 " + c.Request.URL.Path)
		c.SecureJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "网络请求",
			"请求路径":    c.Request.URL.Path,
			"协程":      "同步",
		})
	})

	r.Run(":8080")

}
