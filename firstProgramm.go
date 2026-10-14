package main

import "fmt"

func getVals[K comparable, V any](mp map[K]V) (keys []K, vals []V) {
	for key, val := range mp {
		vals = append(vals, val)
		keys = append(keys, key)
	}

	return
}

type Number interface {
	int | float64 | uint
}

func add[T Number](x T, y T) (r T) {
	r = x + y
	return
}

func main() {
	mp := map[string]int{"a": 100, "b": 200, "c": 300}
	keys, vals := getVals(mp)
	fmt.Println(keys, vals)

	r := add(213, 53)
	fmt.Println(r)
}