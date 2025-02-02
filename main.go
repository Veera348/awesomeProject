package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to our todo list app")
	http.HandleFunc("/hello-go", helloUser)
	http.ListenAndServe("localhost:8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello World"
	fmt.Fprintf(writer, greeting)
}
