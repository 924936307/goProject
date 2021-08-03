package main

import (
	"fmt"
	"strconv"
	"sync"
)

//并发安全的map
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

//fatal error: concurrent map writes
func main1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			key := strconv.Itoa(i)
			set(key, i)
			fmt.Printf("key:%v,value:%v \n", key, i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}

//sync.map
var m1 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			m1.Store(key, i)
			load, ok := m1.Load(key)
			if ok {
				fmt.Printf("k :%v,v:%v \n", key, load)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	count := 0
	m1.Range(func(key, value interface{}) bool {
		count++
		fmt.Println(key, value)
		return true
	})
	//m1中元素个数
	fmt.Println(count)
}
