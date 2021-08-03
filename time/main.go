package main

import (
	"fmt"
	"time"
)

//时间相关的接口

func timestampFormat(timeStamp int64) {
	timeObj := time.Unix(timeStamp, 0)
	fmt.Println(timeObj)
}

func main() {
	now := time.Now()
	fmt.Printf("current time:%v \n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
	fmt.Println("-----------------------获取的时间戳----------------------------")
	millins := now.Unix()
	fmt.Println("当前的秒 ：", millins)
	timestampFormat(1623768811)
	fmt.Println("------------------------测试time.add()----------------------------")
	later := now.Add(time.Hour)
	fmt.Println(later)
	tickDemo()
}

//定时器的使用
func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}
