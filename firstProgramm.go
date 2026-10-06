package main

import (
	"fmt"
	"strconv"
)

// First Function in Go
func main()  {
	x := "1234"
	y, err := strconv.Atoi(x)
	fmt.Printf("The Value of y is %v\n", y)
	if err != nil {
		fmt.Printf("The Error in Conver is %v\n", err)
	} else {
		fmt.Println("Everything Worked Good!!")
	}
}