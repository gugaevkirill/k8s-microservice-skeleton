package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func mainpage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}