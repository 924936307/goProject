package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	//格式化时间
	//12小时制
	str := now.Format("2006-01-02 03:04:05.000 PM Mon Jan")
	fmt.Println(str)
	//24小时制
	format := now.Format("2006-01-02 15:04:05.000 Mon Jan")
	fmt.Println("24小时的格式化", format)
	last := time.Now()
	fmt.Println("执行耗时", last.UnixNano()-now.UnixNano())
	//自定义格式化
	fmt.Println("自定义格式化", now.Format("2006/01/02 03:04:05"))
	fmt.Println("自定义格式化", now.Format("2006-01-02 03:04:05"))
}
