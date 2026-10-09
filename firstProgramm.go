package main

import "fmt"

// First Function in Go
func main()  {
	// Pointer -> arr[0]
	// Length -> 3
	// Capacity -> 5


	arr := [5]int{1, 2 , 3, 4, 5};
	sl := arr[:3]; // [1 2 3], len -> 3, Pointer -> arr[0], Capacity -> 5
	sl = sl[:4] // [1 2 3 4], len -> 4, Pointer -> arr[0], Capacity -> 5
	sl = sl[1:4] // [2 3 4], len -> 3, Pointer -> arr[0], Capacity -> 4
	fmt.Printf("Array: %v\nLength: %v\nCapacity: %v\n", sl, len(sl), cap(sl))
}
