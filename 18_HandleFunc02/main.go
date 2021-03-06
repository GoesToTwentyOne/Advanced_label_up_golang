package main

import (
	"io"
	"net/http"
)

func coding(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Coding Coding Coding")
}

func practice(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Practice Practice Practice")
}

func main() {

	http.Handle("/coding", http.HandlerFunc(coding))
	http.Handle("/practice", http.HandlerFunc(practice))

	http.ListenAndServe(":7777", nil)
}
