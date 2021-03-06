package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nwtl, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conn, err := nwtl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(conn.LocalAddr().String())
	bs := make([]byte, 1024)
	sizeRequest, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	fmt.Println(sizeRequest)
	reqstr := string(bs)
	fmt.Println(reqstr)
	body := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>First test</title>
	</head>
	<body>
	<h1>This is Md.Nihad Hossain</h1>
	</body>
	</html>`

	resp := "HTTP/1.1 200 ok \r\n" + "Content-length : %d\r\n" + "content-type : text/html \r\n" + "\r\n%s"
	// 7 formatting response
	msg := fmt.Sprintf(resp, len(body), body)
	//print response
	fmt.Println(msg)
	// 8 finally send response
	conn.Write([]byte(msg))
	// conn.Close()
	// nwt.Close()

}
