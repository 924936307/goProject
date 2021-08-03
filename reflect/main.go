package main

import (
	"fmt"
	"reflect"
)

//反射
type myInt int64

func reflectType(x interface{}) {
	typeOf := reflect.TypeOf(x)
	fmt.Printf("type :%v,kind:%v \n", typeOf, typeOf.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d \n", int64(v.Int()))
	case reflect.Int32:
		fmt.Printf("type is int32 ,value is %d \n", int32(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f \n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f \n", float64(v.Float()))
	}
}

func main1() {
	var a = 3.14
	reflectType(a)
	var b float32 = 3.14
	reflectType(b)
	var c = 100
	reflectType(c)
	//type name type kind
	var a1 *float32 //指针类型
	var b1 myInt    //自定义类型
	var c1 rune     //类型别名
	reflectType(a1)
	reflectType(b1)
	reflectType(c1)
	type person struct {
		name string
		age  int
	}
	p := person{
		name: "沙河小王子",
		age:  10,
	}
	reflectType(person{})
	reflectType(p)
	//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()的返回值。
	fmt.Println("------------------------------------")
	arr := [...]int{1, 23, 24}
	reflectType(arr) //type :[3]int,kind:array
	s := arr[:]
	reflectType(s) //type :[]int,kind:slice
	d := 100
	reflectType(&d) //type :*int,kind:ptr
	m := make([]int, 12)
	reflectType(m) //type :[]int,kind:slice
	fmt.Printf("-------------华丽的分割线，下面的是value的取值-----------------------")
	g1 := 10
	reflectValue(g1)
	var g2 float64 = 1.23
	reflectValue(g2)
	var g3 float32 = 3.34930439403
	reflectValue(g3)
	var g4 int32 = 11
	reflectValue(g4)
	//查询g1的默认类型
	reflectType(g1)
}

//反射修改变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

//反射中使用专有的Elem()方法来获取指针对应的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)

	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main2() {
	var a int64 = 100
	reflectSetValue1(&a)
	fmt.Printf("修改后的值为：%d \n", a)
	reflectSetValue2(&a)
	fmt.Printf("修改后的值为：%d \n", a)
}

//结构体的反射示例

type stduent struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main3() {
	s := stduent{
		Name:  "bo",
		Score: 10,
	}
	t := reflect.TypeOf(s)
	fmt.Printf("name:%v,kind:%v \n", t.Name(), t.Kind()) //name:stduent,kind:struct
	//获取所有的字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s,index:%d,type:%v,json tag:%v \n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	//通过字段名过去指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s,index:%d,type:%v,json tag:%v \n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

//反射方法的示例
func (s stduent) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s stduent) Sleep() string {
	msg := "hahahah"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("方法的个数：%d \n", t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method:%s \n", methodType)
		name := t.Method(i).Name
		fmt.Printf("method name:%s \n", name)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

//测试反射方法的调用
func main() {
	s := stduent{
		Name:  "bo1",
		Score: 101,
	}
	printMethod(s)
}
