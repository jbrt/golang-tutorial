package main

import (
	"fmt"
)

func main() {
	// Creating a slice without initialization
	var months = make([]string, 3)
	months[0] = "January"
	months[1] = "February"
	months[2] = "March"

	months = append(months, "April")
	fmt.Println(months[3])

	// Creating a slice with initialization
	days := []string{"Saturday", "Sunday"}
	days = append(days, "Monday")
	fmt.Println(days[2])
}
