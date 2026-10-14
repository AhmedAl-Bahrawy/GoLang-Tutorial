package main

import "fmt"

func add[T int | float64 | uint](x T, y T) T {
	return x + y
}



func main() {
	val := add(2, 2.5)
	val2 := add(uint(22), uint(2))
	fmt.Println(val)
	fmt.Println(val2)	
}
