// Package main csv
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
  // csv info
	in := `firstName, lastName, age
  Celina, Jones, 14
  Cailyn, Henderson, 13
  Cayden, Smith, 43
  `
  // read the content of the whole csv
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0])
	}
}
