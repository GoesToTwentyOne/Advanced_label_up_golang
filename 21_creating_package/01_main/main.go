package main

import (
	"fmt"

	calculator "github.com/GoesToTwentyOne/Advanced_label_up_golang/21_creating_package/02_calculator"
)

func main() {
	fmt.Println(calculator.Add(5, 10))
	fmt.Println(calculator.Div(5, 10))
	fmt.Println(calculator.Mul(5, 5))
	fmt.Println(calculator.Sub(40, 5))

}
