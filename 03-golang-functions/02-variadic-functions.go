package main

import (
	"fmt"
)

func variadicSumFunction(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    fmt.Println(variadicSumFunction(1, 3, 5, 7))
}