package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(buffer io.Writer, name string) {
	_, err := fmt.Fprintf(buffer, "Hello, %s", name)
	if err != nil {
		return
	}
}

func HandleMyGreeting(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Chris")
}

func main() {
	err := http.ListenAndServe(":5300", http.HandlerFunc(HandleMyGreeting))
	if err != nil {
		fmt.Println(err)
	}
}
