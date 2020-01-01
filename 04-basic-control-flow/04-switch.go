package main

import (
	"fmt"
)

func main() {
    level := 2
    switch level {
        case 1: 
            fmt.Println("Level 1")
        case 2: 
            fmt.Println("Level 2")
        default:
            fmt.Println("Default choice")
    }
}