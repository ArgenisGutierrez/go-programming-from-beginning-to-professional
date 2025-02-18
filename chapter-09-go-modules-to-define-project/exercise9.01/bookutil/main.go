// Package main creating and using your first module
package main

import "bookutil/author"

func main() {
	// Create an author instance
	authorInstance := author.NewAuthor("Jane Doe", "jane@example.com")
	chapterTitle := "Introduction to Go Modules"
	chapterContenct := "Go modules provide a structured way to manage dependencies and improve code maintainability"
	authorInstance.WriteChapter(chapterTitle, chapterContenct)
	authorInstance.ReviewChapter(chapterTitle, chapterContenct)
	authorInstance.FinalizeChapter(chapterTitle)

}
