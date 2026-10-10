package main

func sum(numbers ...int) (sum int, sum2 int){
	for _, value := range numbers {
		sum += value
		sum2 += value * 2
	}

	return	
}

func main() {
    s, _ := sum(1, 2, 3, 4, 5, 6)
	println(s)

	_, s2 := sum([]int{1, 2, 3, 4, 5, 6}...)
	println(s2)
}
