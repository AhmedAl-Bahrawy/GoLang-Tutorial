package main

import (
	"errors"
	"fmt"
)

func devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Devision by 0")
	}

	return a / b, nil
}

func main() {
	r, err := devide(1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
