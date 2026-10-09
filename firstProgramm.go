package main

import "fmt"

// First Function in Go
func main()  {
	arr := [5]int{1, 2 , 3, 4, 5};
	sl := arr[:3]; // [1 2 3]
	sl2 := arr[1:3]; // [2 3]
	sl[1] = 100
	fmt.Println(sl);
	fmt.Println(sl2);
}
