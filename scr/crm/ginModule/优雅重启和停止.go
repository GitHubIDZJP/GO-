package ginModule

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
manners：优雅关闭的礼貌 Go HTTP 服务器。
graceful：Graceful 是一个 Go 包，可以优雅地关闭 http.Handler 服务器。
grace：Go 服务器的优雅重启和零停机部署。
*/
// 优雅重启或停止
func RrstartOrStop() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		//c.String(http.StatusOK, "欢迎Gin服务器")
		c.SecureJSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"服务请求关闭": "将要执行关闭服务操作",
			"关闭":     "请按control+c",
		})
	})
	//err_fetch := endless.ListenAndServe("localhost:8080", router)
	//print("重新开启服务器:\n", err_fetch)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 引入连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号优雅地关闭服务器
	// 超时5秒.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM--(Kill(没有参数)默认发送syscanll。SIGTERM)
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL但不能被捕获，所以不需要添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("服务器关闭 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//使用 http.Server 内置的 Shutdown() 方法优雅地关机
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.(捕捉ctx.Done()。超时5秒)
	select {
	case <-ctx.Done():
		log.Println("超时5秒.")
	}
	log.Println("服务器退出")

}

/*

在postman测试：get请求后访问为200
在goland编译器终端输入control+c则退出服务器，postman请求则显示Error: connect ECONNREFUSED 127.0.0.1:8080(连接失败),访问Api失败。


输入control+c： ^C2022/11/28 16:41:04 服务器关闭 ...
               超时5秒.
               服务器退出
*/

//新demo
