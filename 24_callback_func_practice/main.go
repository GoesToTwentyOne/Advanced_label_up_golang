package main

import "fmt"

func main() {
	v :=
		func(n []int) int {
			var sum int = 0
			for _, value := range n {
				if value%2 == 0 {
					sum += value
				}

			}
			return sum

		}
	final := goes(v, []int{20, 30, 40, 50, 60, 72})
	fmt.Println(final)
}
func goes(f func(n []int) int, numbers []int) int {
	return f(numbers)

}

// A Higher-Order function is a function that receives a function as an argument or returns
//the function as output.
// Higher order functions are functions that operate on other functions, either by taking them as
//arguments or by returning them.

//callback func
//A "callback func" is when we pass a func into a func as  an argument.
//pass a func into a func as an argument

//https://www.golangprograms.com/
//Tood sir callback func video
//https://www.youtube.com/watch?v=Qsp9htNRRE8&t=361s
