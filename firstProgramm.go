package main

import (
	"fmt"
)

// First Function in Go
func main()  {
	str := "Hello, World!"
	/* str[0] will give the ASCII or UTF-8 format so we need to use string() 
	cuz str[0] will give uint8 */
	// ASCII uses 1 byte
	// UTF-8 uses 4 byte
	fmt.Println(string(str[0]))
}