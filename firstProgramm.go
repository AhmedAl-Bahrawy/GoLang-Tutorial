package main

import "fmt"

// First Function in Go
func main()  {
	mp := map[uint]uint{}
	n := uint(100)

	fmt.Println(mp)

	for number := uint(1); number <= n; number++{
		for d := uint(1); d <= 5; d++ {
			if number % d == 0 {
				mp[d]++
			}
		}
	}

	fmt.Println(mp)
}
