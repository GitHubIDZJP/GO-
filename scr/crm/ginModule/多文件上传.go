package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/*
*
多文件上传就是一次可以上传多个文件，不需要一个文件一个文件上传，这也便于上传文件的人。
获取多文件的MultipartForm方法

MultipartForm：MultipartForm是经过解析的多部分表单（表单里面只有两个属性Value和File），包括文件上传。
File: File部分存储在内存或磁盘上,可通过*FileHeader的Open方法访问。
Value: Value部分存储为字符串。
两者都通过map的字段名进行键控。
*/
func Multiple_file_upload() {
	r := gin.Default()
	r.POST("/multFile", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatal(err)
		}
		//通过字段名映射
		f := form.File["file"]
		//for range遍历文件
		for _, file := range f {
			fmt.Println(file.Filename)
			c.SaveUploadedFile(file, "./"+file.Filename)
			//添加头部映射内容
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s"+file.Filename))
			//File以有效的方式将指定文件写入主体流
			c.File("./" + file.Filename)
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": f[file],
			//})
		}
	})
	r.Run(":8080")

}

/**
接口:
    http://localhost:8080/multFile
两个文件接收并保存下来，并返回给前端了第一个文件

 form-data-要传几张就搞搞几个file参数
*/

/**
FormFile的方法实现
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	//打开请求发送的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//在本地创建一个文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	//把内容拷贝到本地文件
	_, err = io.Copy(out, src)
	return err
}

*/
