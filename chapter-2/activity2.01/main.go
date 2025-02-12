// Package main looping over map data using range
package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	word, count := searchWord(words)
	fmt.Println("Most popular word: ", word)
	fmt.Println("With a count of", count)

}

func searchWord(words map[string]int) (string, int) {
	var (
		w string
		n int
	)
	for word, num := range words {
		if num > n {
			w = word
			n = num
		}
	}
	return w, n
}
