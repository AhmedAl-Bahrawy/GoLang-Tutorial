package main

import "fmt"

// First Function in Go
func main()  {
	// Pointer -> arr[0]
	// Length -> 3
	// Capacity -> 5


	arr := [5]int{1, 2 , 3, 4, 5};
	sl := arr[:3]; // [1 2 3]
	fmt.Printf("Array: %v\nLength: %v\nCapacity: %v\n", sl, len(sl), cap(sl))
}
