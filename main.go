package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

// Run server: go build;
// Test: curl http://127.0.0.1:8000/
func main() {
	router := httprouter.New()
	router.GET("/", mainpage)

	log.Fatal(http.ListenAndServe(":8000", router))
}