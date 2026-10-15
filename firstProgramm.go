package main

import "fmt"


func main() {
	x := 10
	y := &x

	fmt.Printf("%v, %v\n", x, *y)

	*y = 11

	fmt.Printf("%v, %v\n", x, *y)
}