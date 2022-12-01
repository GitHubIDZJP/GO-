package ginModule

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
)

// 加密
func HttpEncrypt() {

	//Example01:---->一行代码启动 LetsEncrypt HTTPS 服务器
	r := gin.Default()
	r.GET("/", func(p *gin.Context) {
		p.String(200, "pong")
	})
	//自定义自动证书管理器示例代码(autotls)
	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))

	//r.Run(":8080")
	/*
		{
		 pong
		}

	*/
}

// 自定义自动证书管理器示例代码
func customAutoCertMgr() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
