package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//request
	nwt, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//connection acccept session start
	conn, err := nwt.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	//remote address print
	fmt.Println(conn.RemoteAddr().String())
	//byte string
	bs := make([]byte, 1024)
	sizeint, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(sizeint)
	//request in string
	reqstr := string(bs)
	fmt.Println(reqstr)
	//response
	body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	<h1>This is Md.Nihad Hossain</h1>
	</body>
	</html>`
	resp := "HTTP/1.1 200 ok \r\n" + "Content-length : %d\r\n" + "content-type : text/html \r\n" + "\r\n%s"
	msg := fmt.Sprintf(resp, len(body), body)
	// fmt.Println(msg)

	conn.Write([]byte(msg))
	// conn.Close()
	// nwt.Close()

}
