// Package main Removing an element from a slice
package main

import "fmt"

func main() {
  words:=[]string{
    "Good",
    "Good",
    "Bad",
    "Good",
    "Good",
  }

  words = append(words[:2],words[3:]... )
  fmt.Println(words)
}
