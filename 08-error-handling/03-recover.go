package main

import (
	"fmt"
)

func functionWillPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover the problem")
		}
	}()

	panic("I got a problem !!")
}

func main() {
	fmt.Println("Before panic")
	functionWillPanic()
	fmt.Println("This string will be printed")
}
