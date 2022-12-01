package ginModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
