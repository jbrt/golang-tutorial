package main

import (
	"fmt"
)

// Outside function without initialization
var myString string

// Outside function with variable initialization
var myHello string = "Hello"

// Variables can be declared as a group of variables
var (
	myWorld string = "World !"
	goodBye string = "Goodbye"
)

func main() {
	myString = "This is a string"
	fmt.Println(myString)
	fmt.Println(myHello)
	fmt.Println(myWorld)
	fmt.Println(goodBye)
}
