package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//原子操作

//对比不同的锁

//接口
type Counter interface {
	Inc()
	Load() int64
}

//普通的结构体
type CommonCounter struct {
	counter int64
}

func (c *CommonCounter) Inc() {
	c.counter++
}

func (c *CommonCounter) Load() int64 {
	return c.counter
}

//拥有互斥锁的结构体
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

//原子锁的结构体
type Atomic struct {
	counter int64
}

func (a *Atomic) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *Atomic) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

//通用的执行函数
func test(c Counter) {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

func main() {
	common := CommonCounter{}
	test(&common)
	a := Atomic{}
	test(&a)
	m := MutexCounter{}
	test(&m)
	fmt.Println(common.Load())
	fmt.Println(a.Load())
	fmt.Println(m.Load())
}
