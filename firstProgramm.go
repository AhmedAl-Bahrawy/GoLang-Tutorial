package main

import "fmt"

func getVals[K comparable, V any](mp map[K]V) (keys []K, vals []V) {
	for key, val := range mp {
		vals = append(vals, val)
		keys = append(keys, key)
	}

	return
}

func main() {
	mp := map[string]int{"a": 100, "b": 200, "c": 300}
	keys, vals := getVals(mp)
	fmt.Println(keys, vals)
}