package main

import (
	"fmt"
)

// Structure declaration
type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street string
	City   string
}

func main() {
	// Structure with initialization
	person := Person{
		Name: "John",
		Age:  32,
		Address: Address{
			Street: "Main street",
			City:   "London",
		},
	}
	// Access to structure member's
	fmt.Println(person.Address.Street)
}
