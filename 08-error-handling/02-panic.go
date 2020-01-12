package main

import (
    "fmt"
)

func main() {
    fmt.Println("Before panic")
    panic("Critical error, can't go ahead...")
    fmt.Println("This string will never be printed")
}