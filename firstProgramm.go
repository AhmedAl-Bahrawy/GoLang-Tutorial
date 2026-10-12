package main

import "fmt"

type Shape interface {
	getParam() uint
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

// End of Triangle Struct

func main() {
	var shapes []Shape = []Shape{Triangle{1, 2, 3}, Sqaure{10}}

	params := uint(0)

	for _, shape := range shapes {
		params += shape.getParam()
	}

	fmt.Println(params)
}
