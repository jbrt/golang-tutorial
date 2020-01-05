package main

import "fmt"

func main() {
	// First version, without initialization
	points := make(map[string]int)
	// Adding one element
	points["John"] = 3
	fmt.Println(points["John"])

	// Second version with initialization
	secondPoints := map[string]int{"John": 3, "Mathew": 5}
	secondPoints["Mark"] = 1
	fmt.Println(secondPoints["Mark"])

	// Deleting one element
	delete(secondPoints, "Mark")
	fmt.Println(secondPoints)
}
