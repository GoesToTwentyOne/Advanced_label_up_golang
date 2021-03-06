package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	netL, err := net.Listen("tcp", "localhost:8888")
	defer netL.Close()
	for {
		if err != nil {
			log.Fatal(err)
		}
		conn, err := netL.Accept()
		if err != nil {
			log.Fatal(err)
			continue

		}
		go handle(conn)

	}
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
		//conn.Write([]byte("I heared you"))
	}
	defer conn.Close()
	fmt.Println("Exit your code")

}
