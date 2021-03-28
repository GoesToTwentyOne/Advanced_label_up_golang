package main

import (
	"fmt"
	"log"
	"os"

	name "github.com/GoesToTwentyOne/Advanced_label_up_golang/22_file_handling/01_file_manipulation/file_create_root_dir"
)

func main() {
	//calling

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}
	fmt.Println(dir)

	//file create with os method
	pointerFile, _ := os.Create("../file_create_root_dir/fun.go")

	nINT, _ := pointerFile.Write([]byte(

		`package name
	 import "fmt"
	func Name(){
		fmt.Println("Hello Gi ")
	}`))
	pointerFile.Close()

	fmt.Println(nINT)
	show()

}

func show() {
	name.Name()
}
