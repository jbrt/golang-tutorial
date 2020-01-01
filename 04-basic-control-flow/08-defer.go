package main

import (
	"fmt"
)

func main() {
    defer fmt.Println("This will be executed after all other statements")
    fmt.Println("Hello")
    fmt.Println("World")
}