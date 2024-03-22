package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	// al ser campos privados, no se puede acceder a ellos
	// var myBook = book.Book{
	// 	title:  "The Catcher in the Rye",
	// 	author: "J.D. Salinger",
	// 	pages:  234,
	// }

	myBook := book.NewBook("The Catcher in the Rye", "J.D. Salinger", 234)
	myBook.PrintInfo()
	myBook.SetTitle("The Catcher in the Rye 2")
	fmt.Println(myBook.GetTitle())

	fmt.Println("=====================")

	myTextBook := book.NewTextBook("Spiderman", "Marvel", 234, "Editorial", "Level")
	myTextBook.PrintInfo()
	fmt.Println(myTextBook.GetTitle())

	fmt.Println("=====================")

	book.PrintBookInfo(myBook)
	book.PrintBookInfo(myTextBook)

	fmt.Println("=====================")

	myDog := animal.Dog{Nombre: "Firulais"}
	myCat := animal.Cat{Nombre: "Garfield"}
	fmt.Println(animal.AnimalSound(&myDog))
	fmt.Println(animal.AnimalSound(&myCat))
}
