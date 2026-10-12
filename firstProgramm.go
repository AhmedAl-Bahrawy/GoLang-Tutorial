package main

import (
	"fmt"
)

type Shape interface {
	getParam() uint
}

// Start of Sqaure Struct
type Sqaure struct {
	Width uint
}

func (s Sqaure) getParam() uint {
	return s.Width * 4
}
// End of Sqaure Struct


func isEvenParam(s Shape) bool {
	return s.getParam() % 2 == 0
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
		if isEvenParam(shape){
			params += shape.getParam()
		}
	}

	fmt.Println(params)
}
