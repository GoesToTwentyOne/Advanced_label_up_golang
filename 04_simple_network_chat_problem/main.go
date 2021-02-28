package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	// bs := make([]byte, 1024)
	// ni, err := conn.Read(bs)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)

	// }
	// fmt.Println(ni)
	// remoteAdddr := conn.RemoteAddr().String()
	fmt.Println(conn.RemoteAddr().String()) //remote address print
	conn.Write([]byte("welcome Nn"))
	conn.Close()
	nl.Close()
}
