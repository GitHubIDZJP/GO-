package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type TestModel struct {
	Name    string
	Age     int
	Title   string
	School  string
	Coupons []ArrCCoupons //优惠券(数组)
}
type ArrCCoupons struct {
	NationalDay       int64 //国庆节
	MidAutumnFestival int   //中秋节
}

func SecureJSONExample() {
	r := gin.Default()

	// 您也可以使用自己的安全json前缀
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		var Name = []string{"刘强", "王维", "周自强", "毛泽东", "周恩来", "任弼时", "刘强东", "欧阳丽丽", "王五", "黄志平"}
		var Title = []string{"锲而不舍", "破釜沉舟", "坚韧不拔", "坚定不移", "事在人为", "投笔从戎", "坚持不懈", "卷土重来", "百折不挠", "晨钟暮鼓"}
		//var ChinaFestival
		arr_a := [10]TestModel{} //定义一个数组长度为10的数组变量arr_a

		var k int
		for k = 0; k < 10; k++ {

			//随机值区间:
			var res_NationalDay int64
			for i := 0; i < 50; i++ {
				res_NationalDay = getRandomWithNo(5, 1000)
				// res := getRandomWithAll(5, 10)
				// res := getRandomWithMin(5, 10)
				// res := getRandomWithMax(5, 10)

			}

			ChinaFestival := ArrCCoupons{
				NationalDay:       res_NationalDay,
				MidAutumnFestival: rand.Intn(100),
			}
			arr_Festival := []ArrCCoupons{ChinaFestival}
			var schoolStr = strconv.Itoa(k)
			SchoolSufficStr := "中"
			SchoolName := schoolStr + SchoolSufficStr
			arr_a[k] = TestModel{
				Name[k],
				rand.Intn(100),
				Title[k],
				SchoolName,
				arr_Festival,
			}
			fmt.Printf("数组下标[%d] = %d\n", k, arr_a[k])
		}

		//names := []string{namesModel, namesModel1, namesModel2}
		// Will output  :   while(1);["lena","austin","foo"]
		//c.SecureJSON(http.StatusOK, names)

		//外层包裹:字典
		Data := map[string]interface{}{
			"status": 1000,
			"data":   arr_a,
		}
		//
		c.SecureJSON(http.StatusOK, gin.H{
			"code":     http.StatusOK,
			"message":  "响应成功",
			"status":   "normal",
			"page":     "1",
			"pageSize": len(arr_a),

			"result": Data,

			"error_code": 0,
		})
		/*
			c.JSON(200, gin.H{

			})
		*/

	})
	r.Run(":8080")
}

// 包含上下限 [min, max]
func getRandomWithAll(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}

// 不包含上限 [min, max)
func getRandomWithMin(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min) + min)
}

// 不包含下限 (min, max]
func getRandomWithMax(min, max int) int64 {
	var res int64
	rand.Seed(time.Now().UnixNano())
Restart:
	res = int64(rand.Intn(max-min+1) + min)
	if res == int64(min) {
		goto Restart
	}
	return res
}

// 都不包含 (min, max)
func getRandomWithNo(min, max int) int64 {
	var res int64
	rand.Seed(time.Now().UnixNano())
Restart:
	res = int64(rand.Intn(max-min) + min)
	if res == int64(min) {
		goto Restart
	}
	return res
}

/*
Api:  http://localhost:8080/someJSON
postman运行效果:
{
    "code": 200,
    "error_code": 0,
    "message": "响应成功",
    "page": "1",
    "pageSize": 10,
    "result": {
        "data": [
            {
                "Name": "刘强",
                "Age": 71,
                "Title": "锲而不舍",
                "School": "0中",
                "Coupons": [
                    {
                        "NationalDay": 6,
                        "MidAutumnFestival": 44
                    }
                ]
            },
            {
                "Name": "王维",
                "Age": 15,
                "Title": "破釜沉舟",
                "School": "1中",
                "Coupons": [
                    {
                        "NationalDay": 6,
                        "MidAutumnFestival": 97
                    }
                ]
            },
            {
                "Name": "周自强",
                "Age": 36,
                "Title": "坚韧不拔",
                "School": "2中",
                "Coupons": [
                    {
                        "NationalDay": 6,
                        "MidAutumnFestival": 11
                    }
                ]
            },
            {
                "Name": "毛泽东",
                "Age": 11,
                "Title": "坚定不移",
                "School": "3中",
                "Coupons": [
                    {
                        "NationalDay": 7,
                        "MidAutumnFestival": 14
                    }
                ]
            },
            {
                "Name": "周恩来",
                "Age": 94,
                "Title": "事在人为",
                "School": "4中",
                "Coupons": [
                    {
                        "NationalDay": 6,
                        "MidAutumnFestival": 9
                    }
                ]
            },
            {
                "Name": "任弼时",
                "Age": 35,
                "Title": "投笔从戎",
                "School": "5中",
                "Coupons": [
                    {
                        "NationalDay": 9,
                        "MidAutumnFestival": 94
                    }
                ]
            },
            {
                "Name": "刘强东",
                "Age": 99,
                "Title": "坚持不懈",
                "School": "6中",
                "Coupons": [
                    {
                        "NationalDay": 9,
                        "MidAutumnFestival": 3
                    }
                ]
            },
            {
                "Name": "欧阳丽丽",
                "Age": 63,
                "Title": "卷土重来",
                "School": "7中",
                "Coupons": [
                    {
                        "NationalDay": 8,
                        "MidAutumnFestival": 29
                    }
                ]
            },
            {
                "Name": "王五",
                "Age": 22,
                "Title": "百折不挠",
                "School": "8中",
                "Coupons": [
                    {
                        "NationalDay": 6,
                        "MidAutumnFestival": 46
                    }
                ]
            },
            {
                "Name": "黄志平",
                "Age": 12,
                "Title": "晨钟暮鼓",
                "School": "9中",
                "Coupons": [
                    {
                        "NationalDay": 6,
                        "MidAutumnFestival": 34
                    }
                ]
            }
        ],
        "status": 1000
    },
    "status": "normal"
}

*/
