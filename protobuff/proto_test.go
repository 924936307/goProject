package protobuff

import (
	"encoding/json"
	"fmt"
	"goProject/protobuff/message/test"
	"log"
	"testing"
	"time"
)

func TestProto(t *testing.T) {
	defer execTime("start exec", time.Now())
	player := &test.PlayerInfo{
		Nickname:     "bobo",
		HeadImageUrl: "url",
		Address:      "广州",
		Gender:       1,
		Privileges:   "躺下",
		Coin:         100,
	}
	//序列化
	//proto.Marshal(player)
	data, err := json.Marshal(player)
	if err != nil {
		fmt.Println("marshal err.")
		return
	}
	newPlayer := &test.PlayerInfo{}
	//反序列化
	//err = proto.Unmarshal(data, newPlayer)
	err = json.Unmarshal(data, newPlayer)
	if err != nil {
		fmt.Println("unmarshal err.")
		return
	}
	fmt.Println(newPlayer)
	//对象是同一个嘛？
	if player.Nickname == newPlayer.Nickname {
		fmt.Println("==============")
	}

	//枚举的使用
	number := test.STATUS(1).Enum().Number()
	fmt.Println("enum number: ", number)
	//test.STATUS_value
	fmt.Println("=== : ", player == newPlayer)
}

//测试对象只声明，是否有初始化值???
var (
	Init test.PlayerInfo
)

func TestInit(t *testing.T) {
	fmt.Println(Init)
	Init.Nickname = "oha"
	Init.Gender = 1
	fmt.Println(Init)
}

func execTime(compareType string, startTime time.Time) {
	consumes := time.Since(startTime).Microseconds()
	log.Printf("exec desc:%s,consumes :%d \n", compareType, consumes)
}
