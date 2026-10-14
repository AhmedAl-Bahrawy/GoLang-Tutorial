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
	defer deffered()

	devision := devided(1, 0)
	fmt.Println(devision)
	
	fmt.Println("Programm Runs Safe")
}
