package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

//测试不带请求参数的get
func TestGet(t *testing.T) {
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("req get error:%v \n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("readall is error:%v \n", err)
		return
	}
	fmt.Print(string(body))
}

//测试get请求带参数
func TestGet2_withParam(t *testing.T) {
	apiUrl := "http://localhost:111/get"
	//url param
	data := url.Values{}
	data.Set("name", "大帅逼")
	data.Set("age", "18")
	url, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("param url requestUrl failed,err: %V,\n", err)
		return
	}
	url.RawQuery = data.Encode()
	fmt.Println(url.String())
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("resp failed,err:%v \n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v \n", err)
		return
	}
	fmt.Println(string(b))
}

//post request with param

func Test_post(t *testing.T) {
	url := "http://localhost:111/post"
	contentType := "application/json"
	data := `{"name":"bobo","age":"12"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post request failed,err:%v \n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("resp failed,err:%v \n", err)
		return
	}
	fmt.Println(string(b))
}
