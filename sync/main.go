package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//测试并发的示例goroutine    go并发模型CSP   goroutime调度 ：Go语言运行时（runtime）层面的实现（GPM）

//go 语言中的并发是通过goroutine实现的
//goroutine类似于线程，属于用户态的线程
//goroutine是由Go语言的运行时（runtime）调度完成
//而线程是由操作系统调度的
//go语言还提供channel给多个goroutine间进行通信
//goroutine 和channel是Go语言秉承CSP并发模型的重要实现基础
//CSP :Communicating Sequential Process

func hello() {
	fmt.Println("hello,goroutine")
}

func main1() {
	go hello()
	fmt.Println("main goroutine done!")
	//main goroutine不暂停等一下，hello()不会输出到控制台
	time.Sleep(time.Second * 2)
}

//sync.WaitGroup来实现goroutine的同步
var wg sync.WaitGroup

func helloGo(i int) {
	defer wg.Done()
	fmt.Println("hello,goroutine!", i)
}

func main3() {
	//设置GPM中p的个数
	//runtime.GOMAXPROCS()
	fmt.Println(runtime.NumCPU())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go helloGo(i)
	}
	wg.Wait() //等待所有登记的goroutine都结束
}

//测试多核心的任务分配情况
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	//先看单线程情况下的执行情况
	/*runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)*/
	///再看2线程情况下的执行情况
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
