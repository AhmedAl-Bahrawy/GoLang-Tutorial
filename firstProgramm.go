package main

import (
	"fmt"
)

// First Function in Go
func main()  {
	str := "Hello, World!"
	
	for _, char := range str {
		fmt.Printf("%c", char)
	}
	
	fmt.Println()
}