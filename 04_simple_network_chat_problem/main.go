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
	remoteADR := conn.RemoteAddr().String()
	fmt.Println(remoteADR)
	conn.Write([]byte("welcome Nn"))
	conn.Close()
	nl.Close()
}
