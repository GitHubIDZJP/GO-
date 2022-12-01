package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 自定义模版--没完全搞懂。有时间得重新复习下
func CustomTemplate() {

	router := gin.Default()
	router.GET("/tmp", func(ctx *gin.Context) {
		html := template.Must(template.ParseFiles("./templates/hello.html", "./templates/hello2.html"))
		//和上面那句一个意思
		//html, _ := template.ParseFiles("./templates/hello.html", "./templates/hello2.html")
		fmt.Println(html.Name())

		ctx.JSON(http.StatusOK, gin.H{
			"html": html,
		})
	})

	//router.SetHTMLTemplate(html)
	router.Run(":8080")

}

/*

ParseFiles是一个标准函数，他对模板文件进行语法分析
ParseFiles函数 只是为了 方便调用Template结构的ParseFiles而设置个一个函数


ParseFiles:方法加载模板
*/

/*
json打印:


{
    "html": {
        "Tree": {
            "Name": "hello.html",
            "ParseName": "hello.html",
            "Root": {
                "NodeType": 11,
                "Pos": 0,
                "Nodes": [
                    {
                        "NodeType": 0,
                        "Pos": 0,
                        "Text": "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9InpoLUNOIj4KPGhlYWQ+CiAgICA8bWV0YSBjaGFyc2V0PSJVVEYtOCI+CiAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoLCBpbml0aWFsLXNjYWxlPTEuMCI+CiAgICA8bWV0YSBodHRwLWVxdWl2PSJYLVVBLUNvbXBhdGlibGUiIGNvbnRlbnQ9ImllPWVkZ2UiPgogICAgPHRpdGxlPkhlbGxvPC90aXRsZT4KPC9oZWFkPgo8Ym9keT4KPHA+SCA="
                    },
                    {
                        "NodeType": 1,
                        "Pos": 250,
                        "Line": 10,
                        "Pipe": {
                            "NodeType": 14,
                            "Pos": 250,
                            "Line": 10,
                            "IsAssign": false,
                            "Decl": null,
                            "Cmds": [
                                {
                                    "NodeType": 4,
                                    "Pos": 250,
                                    "Args": [
                                        {
                                            "NodeType": 5,
                                            "Pos": 250
                                        }
                                    ]
                                }
                            ]
                        }
                    },
                    {
                        "NodeType": 0,
                        "Pos": 253,
                        "Text": "PC9wPgo8L2JvZHk+CjwvaHRtbD4K"
                    }
                ]
            },
            "Mode": 0
        }
    }
}


*/
