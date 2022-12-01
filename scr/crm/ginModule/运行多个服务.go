package ginModule

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

// 第一步加http.Handler
//写到外面类也也行，记得加http.Handler
/*
格式
 func 函数名() http.Handler{

 }

*/

/*
	//-------->直接找到router03.go就看就行，调用的是这个，其实都一样，只是放到外面，通过http.Handler关联起来而已

	func router01() http.Handler {
		e := gin.New()
		e.Use(gin.Recovery()) //第2步：加Recovery
		e.GET("/", func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code":    http.StatusOK,
					"message": "Welcome server 01",
				},
			)
		})

		return e
	}
*/
func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 02",
			},
		)
	})

	return e
}

func RunMultipteServer() {
	/*
			服务1和服务2的区别就是端口号不一样，在postman跑时直接更改端口号就行，就能切换ApiDemo了
		    绑定不一样的port而已
	*/

	/**
	自定义 HTTP 配置里会用到
	s := &http.Server{
			Addr:           ":8080",
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
	*/
	//服务1:http://localhost:8080/
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,  //超时
		WriteTimeout: 10 * time.Second, //超时写入
	}
	//服务2: http://localhost:8081/
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
