package main

import "fmt"

// First Function in Go
func main()  {
	sl := []string{"hello", "world", "hi"}
	test(sl)
	for i, value := range sl {
		fmt.Println(i, value)
	}
}


func test(arr []string) {
	arr[0] = "Changed"
}