package main

import "fmt"


func change(x int) {
	x = 100
}

func pChange(x *int){
	*x = 100
}

func main() {
	a := 10
	change(a)
	fmt.Println(a)

	b := 10
	pChange(&b)
	fmt.Println(b)
}