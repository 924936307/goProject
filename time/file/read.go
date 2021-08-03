package main

import (
	"fmt"
	"io"
	"os"
)

//读取指定文件

func main() {

	file, err := os.Open("../format/format.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer file.Close()
	temp := make([]byte, 12)
	read, err := file.Read(temp)
	if err == io.EOF {
		fmt.Println("文件读完了。")
		return
	}
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Printf("读取了%d个字节的数据", read)
	fmt.Println(string(temp[:]))
}
