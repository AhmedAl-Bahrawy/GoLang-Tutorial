package main

import (
	"fmt"
	"time"
)

func add(x int, y int,ch chan int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println(x + y)
	ch <- x + y
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go add(5, 10, ch, 3)
	go add(10, 15, ch2, 2)
	x := <-ch
	y := <-ch2
	fmt.Println(x, y)
}