package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(dir)
	// err = os.Mkdir("../filemk", 0777)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	//base path
	// basePATH := filepath.Base(dir)
	// fmt.Println(basePATH)
	// //relative path
	// relativePATH := filepath.Join("../filemk")
	// fmt.Println(relativePATH)
	//fullpath
	absulatePATH, _ := filepath.Abs("../filemk")
	// fmt.Println(absulatePATH)
	//make another folder
	// nePATH := filepath.Join(absulatePATH, "../", "../", "04_file")
	// err := os.Mkdir(nePATH, 777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(nePATH)
	newPATH2 := filepath.Join(absulatePATH, "myfile")
	err := os.Mkdir(newPATH2, 777)
	if err != nil {
		log.Fatal(err)
	}

}
