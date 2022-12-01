package ginModule

import "fmt"

// 协程
func InitCoroutine() {

	for i := 0; i < 10; i++ {
		go Add(i, i)
		fmt.Printf("协程值: %d \n", i)
	}

	/*打印:
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	*/
}
func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

/*
协程:
本质上，goroutine 就是协程。 不同的是，Golang 在 runtime、系统调用等多方面对 goroutine 调度进行了封装和处理，当遇到长时间执行或者进行系统调用时。 （推荐学习：go）

会主动把当前 goroutine 的CPU (P) 转让出去，让其他 goroutine 能被调度并执行，也就是 Golang 从语言层面支持了协程。

Golang 的一大特色就是从语言层面原生支持协程，在函数或者方法前面加 go关键字就可创建一个协程。

其他方面的比较

内存消耗方面

每个 goroutine (协程) 默认占用内存远比 Java 、C 的线程少。

goroutine：2KB

线程：8MB

线程和 goroutine 切换调度开销方面

线程/goroutine 切换开销方面，goroutine 远比线程小

线程：涉及模式切换(从用户态切换到内核态)、16个寄存器、PC、SP...等寄存器的刷新等。

goroutine：只有三个寄存器的值修改 - PC / SP / DX.

我们知道，协程（coroutine）是Go语言中的轻量级线程实现，由Go运行时（runtime）管理。
*/

/*

go主线程（有程序员直接称为线程/也可以理解成进程）：一个go线程上，可以起多个协程。可以理解为：协程是轻量级的线程（编译器做优化）

go协程的特点：
1）有独立的栈空间
2）共享程序堆空间
3）调度由用户控制
4）协程是轻量级的线程


*/

/*
1主线程是一个物理线程，直接作用在cpu上。是重量级的，非常耗费cpu资源。
2协程是从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小。
3Golang的协程机制是重要的特点，可以轻松地开启上万个协程。其它编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这里就突显了Golang在并发上的优势


*/
