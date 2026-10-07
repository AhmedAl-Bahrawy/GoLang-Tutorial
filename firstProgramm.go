package main

import "fmt"

// First Function in Go
func main()  {
	var arr [2]int;
	var arr2 [3]int;
	arr3 := [3]int{1, 2, 3}
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	arr5 := [...]int{1, 2, 3, 4, 5, 6}
	arr6 := [...][2]int{{1, 2}, {1, 2}, {1, 2}}

	fmt.Println(arr);
	fmt.Println(arr2);
	fmt.Println(arr3);
	fmt.Println(arr4);
	fmt.Println(arr5);
	fmt.Println(arr6);
}