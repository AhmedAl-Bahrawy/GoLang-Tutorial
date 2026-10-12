package main

import "fmt"

type Shape interface {
	getParam() uint
	getSides() []uint
}


type Sqaure struct {
	width uint
}

func (s Sqaure) getParam() uint {
	return s.width * 4
}

// Start of Triangle Struct
type Triangle struct {
	a uint
	b uint
	c uint
}

func (t Triangle) getParam() uint {
	return t.a + t.b + t.c
}

func (t Triangle) getSides() []uint {
	return []uint{t.a, t.b, t.c}
}
// End of Triangle Struct

func main() {
	var s Shape = Triangle{1, 2, 5}

	fmt.Printf("Param of Shape Triangle is: %v\n", s.getParam())
	fmt.Printf("Sides of Shape Triangle is: %v\n", s.getSides())

}
