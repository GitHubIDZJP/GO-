package grammarLearning

import (
	"os"
)

// 打开本地文件，并且在终端打印文本内容
func OpenLocalFile() {

	buf := make([]byte, 1024)
	// _的意思就是会自动根据后面的类型来判断类型并返回
	f, _ := os.Open("/Users/zoujunping/go文本.txt")
	//defer：推迟
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break

		}
		os.Stdout.Write(buf[:n])
	}
}

// 常量:常量是恒定不变的值 多用于定义程序运行期间不会改变的那些值
// 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值
// const pi  3.14
// const e = 2.17
func Init_constant() {
	//多个常量一起声明
	const (
		age   = 180
		price = 10
	)

	// iota: go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用
	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
	//几个常见的iota示例
	const (
		n11 = iota //0
		n22        //1
		_
		n44 //3
	)
	//声明中间插队
	const (
		n111 = iota //0
		n222 = 100  //100
		n333 = iota //2
		n444        //3
	)
	const n555 = iota //0
}
