package main

import (
	"fmt"
	"io/ioutil"
)

//读取整个文件
func main() {
	file, err := ioutil.ReadFile("../format/format.go")
	if err != nil {
		fmt.Println("read file failed ,err:", err)
		return
	}
	fmt.Println(string(file))
}
