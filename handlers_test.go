package main

import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
)

// Run tests: go test handlers_test.go handler.go

// Check base (/) URL
func TestHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", mainpage)

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expected := "Hi there, I love !"
	actual := strings.Trim(string(response), " \n")
	if actual != expected {
		t.Fatalf(
			"Wrong response '%s', expected '%s'",
			actual, expected,
		)
	}
}
