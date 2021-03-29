package main

import (
	"os"
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
	// absulatePATH, _ := filepath.Abs("../filemk")
	// fmt.Println(absulatePATH)
	//make another folder
	// nePATH := filepath.Join(absulatePATH, "../", "../", "04_file")
	// err := os.Mkdir(nePATH, 777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(nePATH)
	// newPATH2 := filepath.Join(absulatePATH, "myfile")
	// err := os.Mkdir(newPATH2, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	os.Rename(`F:\GOLANG_PROJECTS\src\github.com\GoesToTwentyOne\Advanced_label_up_golang\22_file_handling\03_path_handling_base_relative_aubsulate\filemk\myfile`, `F:\GOLANG_PROJECTS\src\github.com\GoesToTwentyOne\Advanced_label_up_golang\22_file_handling\03_path_handling_base_relative_aubsulate\filemk\myfile2`)

}
