package main

import (
	"fmt"
	"time"
)

func add(x int, y int,ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- x + y
}

func main() {
	ch := make(chan int)
	go add(5, 10, ch)
	x := <-ch
	fmt.Println(x)
}