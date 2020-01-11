package main

import (
	"fmt"
)

// Person structure declaration
type Person struct {
	name string
	age  int
}

// GetName return the name of the person
func (person Person) GetName() string {
	return person.name
}

// SetName can change the name of one person
// Here the receiver is a pointer to the structure.
// If not we can't change the person's name
func (person *Person) SetName(name string) {
	person.name = name
}

func main() {
	// Structure with initialization
	person := Person{
		name: "John",
		age:  32,
	}
	// Access to structure member's
	fmt.Println(person.GetName())

	// Change the name of this person
	person.SetName("William")
	fmt.Println(person.GetName())
}
