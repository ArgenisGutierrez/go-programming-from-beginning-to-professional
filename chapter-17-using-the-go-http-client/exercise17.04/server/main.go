// Package main uploading a file to a web server via a POST request
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upladedFile, uploadedFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	fileContent, err := io.ReadAll(upladedFile)
	if err != nil {
		log.Fatal(err)
	}
	defer upladedFile.Close()

	err = os.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0o600)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("%s Uploaded!", uploadedFileHeader.Filename)))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
