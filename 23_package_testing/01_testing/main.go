package main

import "fmt"

func main() {
	fmt.Println("first testing")
	fmt.Println(Calc(4))

}
func Calc(n int) int {
	result := n + 2
	return result

	// go test
	//go test -v
}
