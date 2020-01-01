package main

import (
	"fmt"
)

func main() {
	var integer int = 123
	var myPointer *int   // Declaration of the pointer
	myPointer = &integer // Storing memory address of integer variable
	fmt.Println(*myPointer)
}
