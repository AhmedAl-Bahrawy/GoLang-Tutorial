package main

import "fmt"

func devided(a int, b int) int {
	return a / b
}

func deffered(){
	fmt.Println("defer")
}

func main() {
	// devision := devided(1, 0)
	// fmt.Println(devision)
	defer deffered()

	fmt.Println("run")
}
