package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func PrintBookInfo(p Printable) {
	p.PrintInfo()
}

// estructura Book con campos privados
type Book struct {
	title  string
	author string
	pages  int
}

// simular un constructor
func NewBook(title, author string, pages int) *Book {
	return &Book{
		title,
		author,
		pages,
	}
}

// "setter" para el campo title
func (b *Book) SetTitle(title string) {
	b.title = title
}

// "getter" para el campo title
func (b *Book) GetTitle() string {
	return b.title
}

// metodo relacionado a la estructura Book
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

// simular un constructor
func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book{
			title,
			author,
			pages,
		},
		editorial,
		level,
	}
}

// metodo relacionado a la estructura Book
func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}
