package main

import (
	"fmt"
	"net/http"
)

func mainpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}