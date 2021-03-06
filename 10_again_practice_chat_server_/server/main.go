package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//lisen browser request
	netL, err := net.Listen("tcp", ":5555")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer netL.Close()
	for {
		//accept network and create a connection
		conn, err := netL.Accept()
		if err != nil {
			fmt.Println(err.Error())

		}
		//print remote address
		fmt.Println(conn.RemoteAddr().String())
		//requst covert on byte
		bs := make([]byte, 1024)
		sizeRequest, err := conn.Read(bs)
		//print request size
		fmt.Println(sizeRequest)
		//request to string
		reqstr := string(bs)
		//print request
		// fmt.Printf("requster message: %s \n", reqstr)
		msg := fmt.Sprintf("your message :%s", reqstr)
		conn.Write([]byte(msg))

	}

}
