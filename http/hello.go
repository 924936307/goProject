package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//  http/hello.tmpl可以识别，./hello.tmpl不能识别
	tmpl, err := template.ParseFiles("http/hello.tmpl")
	if err != nil {
		fmt.Println("create template failed,err: ", err)
		return
	}
	user := UserInfo{
		Name:   "bob",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":666", nil)
	if err != nil {
		fmt.Println("Http server failed,err :", err)
		return
	}
}
