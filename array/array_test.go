package array

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

//测试array

//获取字符串的真实长度
func Test_Array(t *testing.T) {
	str := "Hello,我的天啊！"
	fmt.Println(len(str))
	fmt.Println("the length of str is :", utf8.RuneCountInString(str))
	//或者转换成rune类型的切片
	runes := []rune(str)
	fmt.Printf("the length of \"%s\" is: %d \n", str, len(runes))
	for i, i2 := range str {
		fmt.Println(i, i2)
	}
	fmt.Println("--------------分割线-----------------")
	//数组的测试
	arr := []int{1, 2, 3}
	fmt.Println("array len: ", len(arr))
	for i, i2 := range arr {
		fmt.Println(i, i2)
	}
}
