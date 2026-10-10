package main

func sum(numbers ...int) int{
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum	
}

func main() {
    s := sum(1, 2, 3, 4, 5, 6)
	println(s)
}
