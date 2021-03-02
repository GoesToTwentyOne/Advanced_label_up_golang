package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)
	go ping(c)
	go ping2(c)
	go pong(c)
	go pongp(c)
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Hello")

}
func ping(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "Ping"

	}
}
func ping2(c chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)

	}

}
func pong(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "pong"

	}

}
func pongp(c chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)

	}

}
