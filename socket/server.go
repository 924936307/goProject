package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//https://www.cnblogs.com/liuzhongchao/p/9395027.html
//
func main() {
	address := ":1201"
	tcpAddr, err2 := net.ResolveTCPAddr("tcp4", address)
	check(err2)
	listener, err2 := net.ListenTCP("tcp", tcpAddr)
	check(err2)
	for {
		accept, err2 := listener.Accept()
		if err2 != nil {
			continue
		}
		go handleClient(accept)
	}

}

func handleClient(conn net.Conn) {
	fmt.Println("receive new request")
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		read, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read == 0 {
			break
		} else if strings.TrimSpace(string(request[:read])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			s := time.Now().String()
			conn.Write([]byte(s))
		}
		request = make([]byte, 128) //clear last read content
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
