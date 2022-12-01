package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 设置并获取 cookie
func SetOrFetchCookie() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie值: %s \n", cookie) //test
	})

	router.Run()
}

/*
接口:
http://localhost:8080/cookie


*/
