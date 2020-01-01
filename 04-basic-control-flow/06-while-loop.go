package main

import (
	"fmt"
)

func main() {
	count := 0
	for count < 5 {
		count++
		fmt.Println("Value: ", count)
	}
}
