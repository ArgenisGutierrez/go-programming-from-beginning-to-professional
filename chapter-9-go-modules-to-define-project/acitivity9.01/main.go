// Package main consuming multiple modules
package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	randomQuote := quote.Go()
	fmt.Printf("Generated UUID: %s\n", id)
	fmt.Printf("Random Quote: %s\n", randomQuote)
}
