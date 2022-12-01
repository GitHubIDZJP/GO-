package ginDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接口转为 JSON 格式
// type Student map[string]interface{}

// type Address map[string]interface{ data }

type resultInfo struct {
	stat string

	arrc []zjp_UserInfo //定义数据类型结构体
}

type zjp_UserInfo struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

func MainStudentApi() {

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"}) //设置代理，必须加上，不然下面会提示代理提示权限不安全
	r.GET("/", func(obj *gin.Context) {

		var uData resultInfo
		uData.stat = "sdsd"

		ss := zjp_UserInfo{
			Name:     "王强",
			Age:      "23",
			Title:    "一条倡议 一分钟征来10辆车",
			Category: "头条",
			Url:      "https://mini.eastday.com/mobile/221115053930984942512.html",
		}
		ss1 := zjp_UserInfo{
			Name:     "习近平",
			Age:      "20",
			Title:    "甘肃农信系统切换停业公告",
			Category: "小红书",
			Url:      "https://mini.eastday.com/mobile/221115052848888128599.html",
		}

		vv := []zjp_UserInfo{ss, ss1}
		//fmt.Println("打印数组:\n", uData.arrc)

		Data := map[string]interface{}{
			"stat": "1",
			"data": vv,
		}

		obj.JSON(http.StatusOK, gin.H{

			"status":   200,
			"message":  "响应成功",
			"page":     "1",
			"pageSize": len(vv),
			"result":   Data,
		})

	})
	r.Run(":8080")

}

/*
 {
    "message": "响应成功",
    "page": "1",
    "pageSize": 2,
    "result": {
        "data": [
            {
                "name": "王强",
                "age": "23",
                "title": "一条倡议 一分钟征来10辆车",
                "category": "头条",
                "url": "https://mini.eastday.com/mobile/221115053930984942512.html"
            },
            {
                "name": "习近平",
                "age": "20",
                "title": "甘肃农信系统切换停业公告",
                "category": "小红书",
                "url": "https://mini.eastday.com/mobile/221115052848888128599.html"
            }
        ],
        "stat": "1"
    },
    "status": 200
}


*/
