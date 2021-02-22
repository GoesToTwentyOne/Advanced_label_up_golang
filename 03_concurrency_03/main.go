package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	go say("world")
	go say("hello")
	wg.Wait()

}

func say(s string) {
	for i := 0; i < 5; i++ {

		fmt.Println(s)
		fmt.Println(&s)
		wg.Done()
	}
}
