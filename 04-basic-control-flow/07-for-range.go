package main

import (
	"fmt"
)

func main() {
	data := []int{10, 20, 30, 40}
	for index, value := range data {
		fmt.Println("Index value: ", index, "Data value: ", value)
	}
}
