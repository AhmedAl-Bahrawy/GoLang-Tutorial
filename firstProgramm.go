package main

import (
	"fmt"
	"strings"
)

// First Function in Go
func main()  {
	sl := []string{"Hello", "World"} // Slice Type Of the Array
	numbers := []int{1, 2, 3, 4} // Slice Type Of the Array

	fmt.Println()

	fmt.Printf("Numbers Array Types: %T, numbers: %v, length: %v, capacity: %v\n", numbers, numbers, len(numbers), cap(numbers))
	fmt.Printf("SL Array Types: %T, sl: %v, length: %v, capacity: %v\n", sl, sl, len(sl), cap(sl))

	fmt.Println()
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println()

	for x := 0; x < 10; x++ {
		sl = append(sl, "Ahmed")
		fmt.Printf("SL: %v, length: %v, capacity: %v\n", sl, len(sl), cap(sl))
	}

	fmt.Println()
}
