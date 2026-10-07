package main

import "fmt"

// First Function in Go
func main()  {
	arr := [...][2]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}
	test(arr)
	fmt.Println(arr)
}

func test(arr [4][2]int){
	arr[0] = [2]int{100, 100}
}