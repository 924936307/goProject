package main

import (
	"fmt"
	"strconv"
)

//strconv包实现的基础数据类型和字符串便是的相互转换

func main() {
	//int <==> string
	a := 100000
	itoa := strconv.Itoa(a)
	fmt.Println(itoa)
	atoi, err := strconv.Atoi(itoa)
	if err != nil {
		fmt.Println("can't convert to int")
		return
	}
	fmt.Printf("type:%T,value:%#v \n", atoi, atoi)

	//其他类型的转换
	//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseBool("FALSE"))
	//
}
