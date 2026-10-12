package main

import "fmt"

// Start of  Person Struct
type Sport struct {
	Name string
	Position string
}
// End of Person Struct

// Start of  Person Struct
type Person struct {
	Name string
	Age uint
	FavSport Sport
}
// End of Person Struct



func main() {
	p1 := Person{Age: 17, Name: "Ahmed", FavSport: Sport{Name: "Chess", Position: "IDK"}}
	fmt.Println(p1)
	fmt.Println(p1.FavSport.Name)
}
