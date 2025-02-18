package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	r := rand.Intn(6)
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
