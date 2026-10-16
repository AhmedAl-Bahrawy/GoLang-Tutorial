package main

import (
	"fmt"
	"time"
)



func run() {
	time.Sleep(time.Second * 2)
	fmt.Println("run")
}

func run2() {
	time.Sleep(time.Second * 2)
	fmt.Println("run2")
}

func run3() {
	time.Sleep(time.Second * 2)
	fmt.Println("run3")
}

func main() {
	go run()
	go run2()
	go run3()
	time.Sleep(time.Second * 2)
	fmt.Println("Done")
}