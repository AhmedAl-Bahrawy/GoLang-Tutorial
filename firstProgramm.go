package main

import "fmt"


type Book struct {
	id int
	title string
}

func (b *Book) setTitle(t string) {
	b.title = t
} 

func main() {
	b := Book{id: 1, title: "Mohammed Salah"}
	b.setTitle("Khalid Saqr")

	fmt.Println(b)
}