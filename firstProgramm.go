package main

import "fmt"

// The Passed Functions Must be in the same type
func callFunc(callable func(int) int, number int) int{
	return callable(number)
}

func doubleNumber(number int) int{
	return number * 2;
}

func tripleNumber(number int) int{
	return number * 3;
}

func main()  {
	// Pass Functions to Another Functions
	value := callFunc(doubleNumber, 10)
	fmt.Println(value)

	value = callFunc(tripleNumber, 10)
	fmt.Println(value)
}
