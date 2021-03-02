package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)

}
func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"

	}

}
func printer(c chan string) {

	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 3)
	}
}
