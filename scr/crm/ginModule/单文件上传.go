package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// http://localhost:8080/upload
func A_file_upload() {

	//单文件上传
	router := gin.Default()
	// 为多部分表单设置较低的内存限制 (default is 32 MiB) 默认是32mib
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//上传文件到指定目录
		dst := fmt.Sprintf("./%s", file.Filename)
		// 上传文件到指定的夏令时.
		c.SaveUploadedFile(file, dst) //传到项目里了，项目清单里可以看到postmen请求上传过来的文件 比如 img txt等

		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		c.JSON(http.StatusOK, gin.H{
			"上传信息": dst,
		})
		/*

			接口:  http://127.0.0.1:8080/upload
			{
			    "上传信息": "./CocoaGLPaintSample.png"
			}

		*/
	})
	router.Run(":8080")
	/*
		r := gin.Default()
		r.POST("/FileTest", func(c *gin.Context) {
			//FormFile返回所提供的表单键的第一个文件
			f, _ := c.FormFile("file")
			//SaveUploadedFile上传表单文件到指定的路径
			c.SaveUploadedFile(f, "./"+f.Filename)
			c.JSON(200, gin.H{
				"msg": f,
			})
		})
		r.Run(":8080")
	*/

	/*
				http://127.0.0.1:8080/FileTest
		post请求:
		{
		    "msg": {
		        "Filename": "123.txt",
		        "Header": {
		            "Content-Disposition": [
		                "form-data; name=\"file\"; filename=\"123.txt\""
		            ],
		            "Content-Type": [
		                "text/plain"
		            ]
		        },
		        "Size": 29
		    }
		}


	*/
}
