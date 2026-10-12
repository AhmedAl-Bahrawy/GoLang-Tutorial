package main

import "fmt"

// Start of  Person Struct

type Person struct {
	Name string
	Age uint
}

func (p *Person) getName() string {
	return p.Name
}

func (p *Person) setName(newName string) {
	p.Name = newName
}

// End of Person Struct

func main() {
	p1 := Person{Age: 17}
	p1.setName("Ahmed")
	value := p1.getName()

	fmt.Println(value)
}
