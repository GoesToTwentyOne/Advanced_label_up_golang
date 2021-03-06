package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	netL, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer netL.Close()
	for {
		conn, err := netL.Accept()
		if err != nil {
			log.Fatal(err)
		}
		bs := make([]byte, 1024)
		sizeRequest, e := conn.Read(bs)
		if e != err {
			log.Fatal(e)
		}
		fmt.Println(sizeRequest)
		reqster := string(bs)
		fmt.Println(reqster)
		fmt.Println("your code is run")
	}
}
