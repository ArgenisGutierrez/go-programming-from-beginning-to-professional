// Package main Type checker
package main

import "fmt"

func typeChecker(t any) string {
	switch t := t.(type) {
	case int, int32, int64:
		return fmt.Sprintf("%v is int", t)
	case float32, float64:
		return fmt.Sprintf("%v is float", t)
	case string:
		return fmt.Sprintf("%v is string", t)
	case bool:
		return fmt.Sprintf("%v is bool", t)
	default:
		return fmt.Sprint("{} is unknown")

	}
}

func main() {
  types := []any{
    5,
    5.5,
    "hola",
    false,
  }

  for _, v := range types {
    fmt.Println(typeChecker(v))
  }

}
