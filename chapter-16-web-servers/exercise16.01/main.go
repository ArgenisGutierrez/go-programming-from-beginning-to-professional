// Package main creating a Hello World server
package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hellor World from Go</h1>"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", hello{}))

}
