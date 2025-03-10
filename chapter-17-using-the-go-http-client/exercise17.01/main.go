// Package main sending a GET request to a web server using the Go HTTP Client
package main

import (
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	r, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	log.Println(data)

}
