// Package main using custom headers and options with the Go HTTP Client
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse() string {
	client := http.Client{Timeout: 11 * time.Second}
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "superSecretToken")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
