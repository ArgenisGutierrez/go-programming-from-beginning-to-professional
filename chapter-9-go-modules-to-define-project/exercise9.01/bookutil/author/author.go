package author

import "fmt"

type Author struct {
	Name   string
	Contac string
}

func NewAuthor(name, contact string) *Author {
	return &Author{Name: name, Contac: contact}
}

func (a *Author) WriteChapter(chapterTitle, content string) {
	fmt.Printf("Author %s is writing a chapter titled  '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) ReviewChapter(chapterTitle, content string) {
	fmt.Printf("Author %s is reviewing a chapter titled  '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}
func (a *Author) FinalizeChapter(chapterTitle string) {
	fmt.Printf("Author %s has finalized the chapter titled  '%s'\n", a.Name, chapterTitle)
}
