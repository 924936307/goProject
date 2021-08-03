package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//bufio按行读取

func main() {
	file, err := os.Open("../format/format.go")
	if err != nil {
		fmt.Println("open file failed ,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(readString) != 0 {
				fmt.Println(readString)
			}
			fmt.Println("文件读完了。")
			return
		}
		if err != nil {
			fmt.Println("read file failed ,err:", err)
			return
		}
		fmt.Println(readString)
	}
}
