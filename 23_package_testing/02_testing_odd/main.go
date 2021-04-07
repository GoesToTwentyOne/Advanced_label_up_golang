package main

import "fmt"

func main() {
	fmt.Println(Odd(3))

}
func Odd(n int) int {
	if n%2 == 0 {
		return 0

	}
	return 1
}
