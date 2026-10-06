package main

import (
	"fmt"
)

// First Function in Go
func main()  {
	a := 2

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough
	default:
		fmt.Println("default")
	}
}