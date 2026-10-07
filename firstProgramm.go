package main

import (
	"fmt"
)

// First Function in Go
func main()  {
	str := "Hello, World!"
	
	for idx := 0; idx < len(str); idx++ {
		if string(str[idx]) == string(str[len(str) -1 ]){
			fmt.Printf("%c\n", str[idx])
		} else {
			fmt.Printf("%c", str[idx])
		}
	}
}