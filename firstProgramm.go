package main

import (
	"fmt"
	"time"
)

func add(x int, y int,ch chan<- int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	ch <- x + y
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go add(5, 10, ch, 2)
	go add(10, 15, ch2, 2)

	for i := 0; i < 2; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		}
	} 
	
}