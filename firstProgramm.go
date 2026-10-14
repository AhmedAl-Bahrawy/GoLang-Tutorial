package main

import "fmt"

func devided(a int, b int) int {
	return a / b
}

func deffered(){
	r := recover()
	if r == nil {
		fmt.Println("No Err")
	} else {
		fmt.Println(r)
	}
}

func main() {
	// devision := devided(1, 0)
	// fmt.Println(devision)
	fmt.Println("run-1")
	defer deffered()

	// panic("Here We Are There is an Err Here Fix that Bro")

	fmt.Println("run-2")
}
