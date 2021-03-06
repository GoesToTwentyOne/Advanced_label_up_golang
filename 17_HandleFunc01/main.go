package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/contact", c)
	http.HandleFunc("/about", a)
	http.ListenAndServe(":8080", nil)
}
func h(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, ` This is home page `)

}
func c(response http.ResponseWriter, req *http.Request) {
	io.WriteString(response, "This is contact page")
}
func a(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "This is about page")
}
