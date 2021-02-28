package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//==>Requst send web browser
	// 1 request listen
	nwt, err := net.Listen("tcp", ":8888")
	//error handling
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// 2 connection acccept session start
	conn, err := nwt.Accept()
	//error handling
	if err != nil {
		fmt.Println(err.Error())
	}
	//remote address print
	fmt.Println(conn.RemoteAddr().String())
	//byte string
	bs := make([]byte, 1024)
	// 3 byte string received and pass
	sizeint, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
	}
	//byte string size
	fmt.Println(sizeint)
	//4 request in conversion string
	reqstr := string(bs)
	//web browser request  print
	fmt.Println(reqstr)
	//5 write response body
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
	// 6 response field must be write
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
