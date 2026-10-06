package main

import (
	"fmt"
	"math"
)

// First Function in Go
func main()  {
	x := float64(10)
	y := float64(20)

	fmt.Println(math.Min(x, y))
	fmt.Println(math.Max(x, y))
}