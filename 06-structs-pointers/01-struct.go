package main

import (
	"fmt"
)

// Structure declaration
type Person struct {
    Name string
    Age  int
}

func main() {
    // Structure with initialization
    person := Person{Name: "John", Age: 32}
    // Access to structure member's
    fmt.Println(person.Name)
}