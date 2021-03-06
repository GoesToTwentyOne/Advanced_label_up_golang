package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	netL, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal(err)
	}
	defer netL.Close()
	for {
		conn, err := netL.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, "\nHello from tcp server\n")
		fmt.Fprintf(conn, "How is your day!\n")
		fmt.Fprintf(conn, "well")
		conn.Close()

	}
}
