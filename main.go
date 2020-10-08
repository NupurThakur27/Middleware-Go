package main

import (
	"fmt"
	"net/http"
	"time"
)

func now() string {
	return time.Now().Format(time.Stamp) + " "
}

//Middleware
func logger(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler started at: " + now())
		defer fmt.Println("Handler ended at: " + now())
		f(w, r)
	}
}

//Basic Handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	fmt.Fprintf(w, "Check your Terminal")
}

func main() {

	http.HandleFunc("/", logger(handler))
	http.ListenAndServe(":8080", nil)
}
