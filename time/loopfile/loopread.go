package main

import (
	"fmt"
	"io"
	"os"
)

//循环读取文件
func main() {
	file, err := os.Open("../format/format.go")
	if err != nil {
		fmt.Println("open file failed ,err:", err)
		return
	}
	defer file.Close()
	//循环读取文件
	//定义一个切片
	var content []byte
	var temp = make([]byte, 128)
	for {
		n, err := file.Read(temp)
		if err == io.EOF {
			fmt.Println("文件读完了。")
			break
		}
		if err != nil {
			fmt.Println("read file failed ,err:", err)
			return
		}
		fmt.Println("本次读取的长度：", n)
		//一个一个元素的添加
		content = append(content, temp[:]...)
	}
	fmt.Println(string(content))
}
