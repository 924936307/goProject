package main

import (
	"fmt"
	_ "golang.org/x/sys/unix"
	"runtime"
	"time"
)

func main() {
	//fmt.Println(runtime.NumCPU())
	//使用一个内核可以保证执行顺序
	runtime.GOMAXPROCS(1)
	fmt.Println("测试开始。。。。。")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	fmt.Println("----------------想要的结果------------------------")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("----------------实现想要的结果-------------------------")
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
		go func() {
			fmt.Println(<-ch)
		}()
	}
	time.Sleep(3 * time.Second)
}
