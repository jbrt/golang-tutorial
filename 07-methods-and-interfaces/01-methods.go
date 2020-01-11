package main

import (
	"fmt"
)

// Structure declaration
type Person struct {
	name string
	age  int
}

func (person Person) GetName() string {
	return person.name
}

func main() {
	// Structure with initialization
	person := Person{
		name: "John",
		age:  32,
	}
	// Access to structure member's
	fmt.Println(person.GetName())
}
