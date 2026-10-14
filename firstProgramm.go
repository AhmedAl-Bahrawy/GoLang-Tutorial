package main

import "fmt"

func add[T int | float64 | uint](x T, y T) (sum T) {
	sum = x + y
	return
}



func main() {
	val := add(2, 2.5)
	val2 := add(uint(22), uint(2))
	fmt.Println(val)
	fmt.Println(val2)	
}
