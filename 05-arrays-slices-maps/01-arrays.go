package main

import (
	"fmt"
)

func main() {
	// First kind of declaration
	var listOfDays [2]string
	listOfDays[0] = "Sunday"
	listOfDays[1] = "Monday"
	fmt.Println(listOfDays[1])

	// second kind of declaration
	listOfMonths := []string{"March", "April"}
	fmt.Println(listOfMonths[0])
}
