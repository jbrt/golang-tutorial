package main

import (
	"errors"
	"fmt"
)

func returnAnErrorIfValuesNotEqual(first, second int) (bool, error) {
	if first != second {
		err := errors.New("These numbers are not equals")
		return false, err
	}
	return true, nil
}

func main() {
	_, err := returnAnErrorIfValuesNotEqual(1, 2)
	if err != nil {
		fmt.Println("Error while calling function: ")
		fmt.Println(err)
	}
}
