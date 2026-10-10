package main

import "fmt"

type Person struct {
	name string
	age uint
	f func(string) string
}

func getName(p Person) string {
	return p.name
}

func main() {
	var p1 Person = Person{name: "Ahmed"}
	p1.name = "Ahmed Albahrawy"
	p1.f = func(x string) string {
		return x + "s"
	}

	
	fmt.Printf("name: %v, age: %v\n", p1.name, p1.age)

	fmt.Println(getName(p1))
}
