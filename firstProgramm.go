package main

import "fmt"

// First Function in Go
func main()  {
	var mp map[string]int = map[string]int{"a" : 1}
	mp2 := map[string]int{"b" : 2}
	mp3 := make(map[string]int)

	// value type is a slice of int
	mp4 := map[string][]int{"c" : {1, 2, 3}}
	mp4["b"] = []int{1, 2 ,3}
	mp4["d"] = []int{1, 2, 3}
	delete(mp4, "b")

	// OK is to check if the key is exists
	value, ok := mp4["b"]
	fmt.Println(value, ok)



	fmt.Printf("mp: %v\n", mp)
	fmt.Printf("mp2: %v\n", mp2)
	fmt.Printf("mp3: %v\n", mp3)
	fmt.Printf("mp4: %v\n", mp4)
}
