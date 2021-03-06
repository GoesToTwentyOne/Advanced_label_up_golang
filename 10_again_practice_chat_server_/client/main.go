package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":5555")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conn.Write([]byte("\nHello\n"))
}
