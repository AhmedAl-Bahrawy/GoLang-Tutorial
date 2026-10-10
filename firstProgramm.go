package main

import "fmt"

// First Function in Go
func main()  {
	var mp map[string]int = map[string]int{"a" : 1}
	mp2 := map[string]int{"b" : 2}
	mp3 := make(map[string]int)

	fmt.Printf("%v\n", mp)
	fmt.Printf("%v\n", mp2)
	fmt.Printf("%v\n", mp3)
}
