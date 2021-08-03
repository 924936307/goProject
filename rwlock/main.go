package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//读写锁的测试类
var (
	x      int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
	lock   sync.Mutex
)

func write() {
	rwlock.Lock() //添加写锁
	x = x + 1
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func read() {
	rwlock.RLock() //添加读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() //解读锁
	wg.Done()
}

func main() {
	fmt.Println(runtime.NumCPU())
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
	fmt.Println(x)
}
