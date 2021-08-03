package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "AHTGQTcI5cmP7n4l",
		DB:           1,
		PoolSize:     100,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Microsecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})
	str, error := Client.Ping().Result()
	if error != nil {
		panic("init redis error")
	} else {
		fmt.Println("init redis ok")
	}
	fmt.Println("init succeed return: ", str)
}

func get(key string) (string, bool) {
	ret, error := Client.Get(key).Result()
	if error != nil {
		return "", false
	} else {
		return ret, true
	}
}

func set(key string, val string, expTime int32) {
	Client.Set(key, val, time.Duration(expTime)*time.Second)
}

func main() {
	fmt.Println("测试redis的基本用法。")
	set("test", "111", 1000)
	str, ok := get("test")
	if !ok {
		fmt.Println("get error")
	} else {
		println(str)
	}
}
