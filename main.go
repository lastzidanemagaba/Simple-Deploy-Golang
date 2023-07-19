package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", SayHello)
	http.ListenAndServe(":8989", handler)
	//we tell our api to listen to all request to port 8080.
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello world`)
}

func SaySomething(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello world`)
}
