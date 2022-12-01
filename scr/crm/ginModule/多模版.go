package ginModule

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

// 下载并安装 go get -u github.com/gin-contrib/multitemplate
func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("webPageTemplate", "templates/webPageTemplate.html", "templates/webPageTemplate.html")
	//	r.AddFromFiles("article", "templates/Qatar.html", "templates/index.html", "templates/article.html")

	r.AddFromFiles("article", "templates/Qatar.html", "templates/order.html", "templates/webPageTemplate.html")
	return r
}
func MultipleTemplatePlate() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "webPageTemplate", gin.H{
			"htmlTemplate": "Html5模板引擎---->邹俊平",
		})
	})
	/**
	Api：
	http://localhost:8080/

	<body>
		模板类型:Html5 Template Engine
	</body>



	/通过{{.htmlTemplate}}里的进行.htmlTemplate关联
	*/

	//2
	/*router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "html标注",
			"name":  "习近平",
			"age":   "20岁",
		})

	})*/
	/*
	   接口: http://localhost:8080/article
	      <body>
	      	html标注
	      	习近平
	      	20岁
	      </body>
	*/
	/**
	//3
		router.GET("/article", func(c *gin.Context) {
			c.HTML(200, "webPageTemplate", gin.H{
				"htmlTemplate": "Html5模板引擎---->邹俊平",
				//"title": "html标注",
				//"name":  "习近平",
				//"age":   "20岁",
			})
	*/
	/*
		接口: http://localhost:8080/article

			<body>
				模板类型:Html5模板引擎----&gt;邹俊平
			</body>


	*/
	//4
	var files []string
	filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	router.LoadHTMLFiles(files...)
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "order.html", gin.H{
			"age":   10,
			"name":  "刘清政",
			"hobby": [3]string{"抽烟", "喝酒", "烫头"},
			"wife":  []string{"刘亦菲", "迪丽热巴", "古力娜扎"},
			"info":  map[string]interface{}{"height": 180, "gender": "男"},
			"book":  Book{"红楼梦", 99},
		})

	})
	router.Run(":8080")

	/*
				打印:
		接口:http://localhost:8080/article
				<body>
					<h1>渲染字符串，数字，数组，切片，maps，结构体</h1>
					<p>年龄：10</p>
					<p>姓名：刘清政</p>
					<p>爱好：[抽烟 喝酒 烫头]</p>
					<p>wife：[刘亦菲 迪丽热巴 古力娜扎]</p>
					<p>信息：map[gender:男 height:180]--->男</p>
					<p>图书：{红楼梦 99}--->红楼梦</p>
				</body>
	*/
}

type Book struct {
	Name  string
	price int
}
