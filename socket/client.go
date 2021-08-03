package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = tcpConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result := make([]byte, 128)
	for {
		read, err := tcpConn.Read(result)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println("client receive resp :", string(result[:read]))
		break
	}
	//ioutil.ReadAll会阻塞
	/*bytearr, err := ioutil.ReadAll(tcpConn)
	checkError(err)
	fmt.Println("client receive resp: ",string(bytearr))*/
	//close
	//tcpConn.Close()
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return
	}
}
