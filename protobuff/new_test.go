package protobuff

import (
	"fmt"
	"testing"
)

//测试new()的对象和直接:=声明的对象的区别
//new()出来的对象是指针类型
type student struct {
	name   string
	age    int8
	gender int8
}

func Test_new(t *testing.T) {
	stu := student{
		name:   "bobo",
		age:    12,
		gender: 1,
	}
	stu1 := new(student)
	//different : stu :protobuff.student,   stu1: *protobuff.student
	fmt.Printf("different : stu :%T,stu1: %T \n", stu, stu1)
	fmt.Println(stu.age)
	fmt.Println(stu1.age)
	//测试声明和初始化是否是一起的
	stu_null := student{}
	fmt.Println(stu_null.name)
	fmt.Println(stu_null.gender)
	fmt.Println(stu1.name)
	fmt.Println(stu1.gender)
}
