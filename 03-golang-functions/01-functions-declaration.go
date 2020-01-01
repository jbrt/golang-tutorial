package main

import (
	"fmt"
)

func withoutArgsAndNoReturnValue() {
	fmt.Println("No arguments and no return values")
}

func withArgsAndNoReturnValue(alpha int, beta int) {
	fmt.Println(alpha + beta)
}

func withArgsAndReturnValue(alpha int, beta int) int {
	return alpha + beta
}

func withArgsAndMultipleReturnValue(alpha int, beta int) (int, int) {
	return alpha * 2, beta * 2
}

func main() {
	withoutArgsAndNoReturnValue()
	withArgsAndNoReturnValue(100, 100)
	fmt.Println(withArgsAndReturnValue(200, 300))
	a, b := withArgsAndMultipleReturnValue(400, 400)
	fmt.Println(a, b)
}
