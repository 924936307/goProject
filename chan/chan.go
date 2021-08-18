package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//channel 通道
//chan示例  实现数据再函数间的交换

//go语言的并发模型是CSP,提倡通过通信共享内存而不是通过共享内存而实现通信
func main1() {
	//创建通道
	//声明
	var ch chan int
	fmt.Println(ch) //nil
	//分配空间
	ch = make(chan int, 1)
	fmt.Println(ch)

	ch1 := make(chan []int)
	fmt.Println(ch1)
	//赋值
	ch <- 10
	//接收通道中的值
	x := <-ch
	//接收值，忽略结果
	<-ch
	fmt.Println(x)
	//关闭通道
	close(ch)
}

//测试有缓冲区的通道
func recv(c chan int) {
	ret := <-c
	fmt.Println(ret)
}

func main2() {
	ch := make(chan int, 1)
	ch <- 10
	go recv(ch)
	time.Sleep(time.Second)
}

//无缓存的通道测试
//如何判断一个通道是否关闭
func main3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		fmt.Println("开始执行ch1")
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		//通道关闭后再插入会引发panic
		close(ch1)
		fmt.Println("ch1执行完毕")
	}()

	go func() {
		fmt.Println("开始执行ch2")
		for i := 0; i < 100; i++ {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
		fmt.Println("ch2执行完毕")
	}()
	//当通道被关闭的时候就会退出for range
	for i := range ch2 {
		fmt.Println(i)
	}
	fmt.Println("main goroutine执行完毕。")
}

//worker pool
//chan<-    <-chan
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d,start job: %d \n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d,end job: %d \n", id, job)
		results <- job * 2
	}
}

func main4() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	time.Sleep(time.Second * 20)
	//close(results)
	for result := range results {
		fmt.Println(result)
	}
}

//select 多路复用
func main5() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

//互斥锁/读写锁
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func add() {
	for i := 0; i < 100; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main6() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

var ch1 chan int

func fuzhi() {
	for i := 0; i < 10; i++ {
		ch1 <- i
		fmt.Println("写入chan.:", i)
	}
	//close(ch1)
}

func main() {
	ch1 = make(chan int, 20)
	go fuzhi()
	//在main goroutine线，期望从管道中获得一个数据，而这个数据必须是其他goroutine线放入管道的
	//但是其他goroutine线都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。
	//所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。
	//这显然是没有结果的，所以这个程序就说“算了吧，不坚持了，我自己自杀掉，报一个错给代码作者，我被deadlock了”
	//解决办法是close()或者用其他的子goroutine来遍历chan
	/*for{
	if i, ok := <-ch1; ok {
		fmt.Println(i)
	}else {
		break
	}
	}*/
	go func() {
		for i := range ch1 {
			fmt.Println("从通道中读取的参数", i)
		}
	}()
	for {
		runtime.GC()
	}
}
