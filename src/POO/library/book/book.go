package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	// Tittle string//Publico
	// Author string//Publico
	// Pages  int//Publico
	tittle string //privado
	author string //privado
	pages  int    //privado
}

//Constructor
func NewBook(tittle, author string, pages int) *Book { // Sirve para cuando las propiedades sean privadas
	return &Book{
		tittle: tittle,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTittle(tittle string) {
	b.tittle = tittle
}

func (b *Book) Gettittle() string {
	return b.tittle
}

func (b *Book) PrintInfo() {
	fmt.Printf("Tittle: %s\nAuthor: %s\nPages: %d\n", b.tittle, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(tittle, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{tittle, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Tittle: %s\nAuthor: %s\nPages: %d\nEditioral: %s\nLevel: %s\n", b.tittle, b.author, b.pages, b.editorial, b.level)
}
