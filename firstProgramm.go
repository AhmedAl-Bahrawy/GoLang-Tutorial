package main

import (
	"fmt"
)

// First Function in Go
func main()  {

	// Use For Loop as For
	for idx := 0; idx < 10; idx++ {
		fmt.Printf("idx: %v\n", idx);
	}

	// Use For Loop as While
	x := 0
	for x < 10 {
		fmt.Printf("x: %v\n", x);
		x++;
	}

	// Use For Loop as While True
	y := 0
	for true {
		fmt.Printf("y: %v\n", y);
		y++;
	}
}