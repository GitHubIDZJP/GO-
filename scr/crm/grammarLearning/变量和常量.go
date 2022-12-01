package grammarLearning

import "fmt"

// 全局变量
var m = 100

func Init_variable() {

	//	var 变量名 变量类型

	//声明变量关键字以var开头，变量类型放在变量的后面，行尾无需分号
	/*
		    var name String
			var age init
			var isOk bool
	*/
	//批量声明
	//每声明一个变量就需要写var关键字会比较繁琐，go语言中还支持批量变量声明
	/*
		var (
			a string
			b int
			c bool
			d float32
		)
	*/

	//声明变量的时候为其指定初始值
	// var  变量名 类型 = 表达式
	// var name string = "pprof.cn"
	// var sex int = 1

	// 一次初始化多个变量
	//var name, sex = "sdds", 2
	//fmt.Println("name", "sex")

	//类型推导
	n := 10
	m := 200 //声明局部变量m
	fmt.Println(m, n)
	//回调
	x, _ := foo()
	_, y := foo()
	/*注意事项：
	函数外的每个语句都必须以关键字开始（var、const、func等）
	    :=不能使用在函数外。
	    _多用于占位，表示忽略值
	*/
	fmt.Printf("x=", x)
	fmt.Printf("y=", y)
}

// 匿名变量
func foo() (int, string) {
	return 10, "123"
}
