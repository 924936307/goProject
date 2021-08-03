package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func gethandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	//请求类型是application/json时从r.body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request,body failed,err:%v \n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status":"post success"}`
	w.Write([]byte(answer))
}

//带参数的get请求的服务端测试
func Test_server(t *testing.T) {
	http.HandleFunc("/get", gethandler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe("localhost:111", nil)
}
