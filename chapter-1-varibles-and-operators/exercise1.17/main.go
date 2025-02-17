// Package main Constants
package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit
const (
	CacheKeyBook = "book_"
	CackeKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func getBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func setBook(isbn, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func getCD(sku string) string {
	return cacheGet(CackeKeyCD + sku)
}

func setCD(sku, title string) {
	cacheSet(CackeKeyCD+sku, title)
}

func main() {
	cache = make(map[string]string)
	setBook("1234-5678", "Get Ready to go")
	setCD("1234-5678", "Get Ready to go Audiobook")
	fmt.Println("Book: ", getBook("1234-5678"))
	fmt.Println("CD: ", getCD("1234-5678"))
}
