package main

import (
	"fmt"
	"math"
)

// First Function in Go
func main()  {
	x := float64(10)
	y := float64(20)

	fmt.Printf("The Min Value Between X and Y is: %v\n", math.Min(x, y))
	fmt.Printf("The Max Value Between X and Y is: %v\n", math.Max(x, y))
	fmt.Printf("The Power 3 of the x is: %v\n", math.Pow(x, 3))
}