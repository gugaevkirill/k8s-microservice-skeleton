package main

import (
	"net/http"
)

// Run server: go build;
// Test: curl http://127.0.0.1:8000/
func main() {
		http.HandleFunc("/", mainpage)
		http.ListenAndServe(":8000", nil)
}