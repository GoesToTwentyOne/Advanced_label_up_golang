package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("what do you want Encoding or Decoding if you need Encoding press 1 or press 2")
	input := bufio.NewReader(os.Stdin)
	fmt.Println("enter your choice 1=Encoding 2=Decoding :")
	option, _ := input.ReadString('\n')
	optionRating, _ := strconv.ParseInt(strings.TrimSpace(option), 10, 64)

	if optionRating == 1 {
		fmt.Println("enter your email :")
		email, _ := input.ReadString('\n')
		fmt.Println("enter your password :")
		password, _ := input.ReadString('\n')

		encription(email, password)

	} else {
		fmt.Println("enter your data :")
		email, _ := input.ReadString('\n')
		decription(email)

	}

}
func encription(s, s1 string) {
	resp := fmt.Sprintf("\x00"+" %s"+"\x00"+"%s", s, s1)
	smtpAuthENC := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println("Your encription data : ", smtpAuthENC)

}
func decription(d string) {
	data, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Your decription data : ", string(data))

}
