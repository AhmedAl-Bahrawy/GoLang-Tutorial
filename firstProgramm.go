package main

import "fmt"

// First Function in Go
func main()  {
	fmt.Println("Hello, World!");

	var x string = "X";
	fmt.Printf("%T", x);
	fmt.Println();

	y := int(10)
	fmt.Printf("%T", y)
	fmt.Println()

	z := uint(y)
	fmt.Printf("%T", z)
	fmt.Println()
}