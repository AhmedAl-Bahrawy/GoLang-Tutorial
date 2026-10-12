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
	FavSports []Sport
}
// End of Person Struct



func main() {
	p1 := Person{Age: 17, Name: "Ahmed",
		FavSports: []Sport{
			Sport{"Chess", "HAMADA"},
			Sport{"Football", "CR7"}}}

	
	fmt.Println(p1)
	fmt.Println(p1.FavSports[0].Name)
	fmt.Println(p1.FavSports[1].Name)
}
