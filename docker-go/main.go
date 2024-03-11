package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Go in docker 🐳")
}