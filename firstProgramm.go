package main

import "fmt"

// First Function in Go
func main()  {
	arr := [...][2]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}

	arr[0] = [2]int{3, 4}
	
	for _, nested := range arr {
		for _, value := range nested {
			fmt.Println(value)
		}
	}
}