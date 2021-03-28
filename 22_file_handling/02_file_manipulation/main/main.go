package main

import (
	"fmt"
	"log"
	"os"

	print "github.com/GoesToTwentyOne/Advanced_label_up_golang/22_file_handling/02_file_manipulation/filedir"
)

func main() {
	fmt.Println(createFile("print.go",
		`package print
		import "fmt"
	func Goes(){
		fmt.Println("my name is Rowja")
		
	}
	`))
	pri()

}
func createFile(filename, filecode string) bool {
	pointerFile, err := os.Create("../filedir/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer pointerFile.Close()
	_, err = pointerFile.Write([]byte(filecode))
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func pri() {
	print.Goes()

}
