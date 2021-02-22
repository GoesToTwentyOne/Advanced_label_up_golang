package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func main() {
	waitgroup.Add(2)

	go nihad()
	go rowjatul()
	waitgroup.Wait()

}
func nihad() {
	fmt.Println("I am Nihad")
	waitgroup.Done()
}
func rowjatul() {
	fmt.Println("I am Rowjatul")
	waitgroup.Done()
}
