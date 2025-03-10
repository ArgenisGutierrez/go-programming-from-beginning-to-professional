// Package main a stylish welcome
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
  http.Handle(
    "/public/",
    http.StripPrefix(
      "/public/",
      http.FileServer(http.Dir("./public")),
    ),
  )
	log.Fatal(http.ListenAndServe(":8080", nil))
}
