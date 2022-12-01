package ginModule

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// 预订包含绑定和验证的数据
type Booking struct {
	//预订日期回传字段
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		//获取现在日期时间
		today := time.Now()
		//如果今天的未来日期是当前日期，也就是未来日期CheckOut !> CheckIn则返回false
		if today.After(date) {
			return false
		}
	}
	return true
}

// 自定义验证器(也可以注册自定义验证器）
func MainCustomValidators() {
	route := gin.Default()

	/*
		binding.绑定
		Validator.验证
		Engine引擎
	*/
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	//可预订的(bookable）
	route.GET("/bookable", getBookable)
	route.Run(":8080")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "预订日期有效!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

/**
get请求
接口: http://localhost:8080/bookable
打印:
{
    "error": "Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'required' tag\nKey: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'required' tag"
}
报错解析:没有回传2和必须字段check_in，check_out所以报错
*/

/**
接口:
http://localhost:8080/bookable?check_in=2030-04-16&check_out=2070-04-19
get请求
打印:
{
    "message": "预订日期有效!"
}

*/
/*
get请求:
http://localhost:8080/bookable?check_in=2030-04-16&check_out=2010-04-19
打印:
{
    "error": "Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"
}
报错解析:
回传字段值日期只能check_in日期必须小于check_out，反之则报错

*/
