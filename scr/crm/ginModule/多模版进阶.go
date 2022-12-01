package ginModule

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"time"
)

// 多模版进阶

func MultipleTemplateAdvanced() {
	router := gin.Default()
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "discipline.html", gin.H{
			"subjects": "物理!",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article.html", gin.H{
			"weight":     "120斤",
			"iotProject": "智能家居",
		})
	})
	/*
		接口: http://localhost:8080/article
		<body>
			120斤
			智能家居
		</body>

	*/
	var fetchTime = cutDate(time.Now())
	fmt.Println("时间:\n", fetchTime)
	router.Run(":8080")

}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	//生成模板映射 layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		// 加模板函数的加载方法 r.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
func cutDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
