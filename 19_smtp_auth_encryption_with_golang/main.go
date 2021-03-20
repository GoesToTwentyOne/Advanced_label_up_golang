package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	resp := []byte("\x00" + "nihadhossain22@gmail.com" + "\x00" + "AlexGoot210@")
	fmt.Println(string(resp), resp)
	// smtpAuthENC := base64.StdEncoding.EncodeToString([]byte(resp))
	smtpAuthENC := base64.StdEncoding.EncodeToString(resp)
	fmt.Println(smtpAuthENC)

}
