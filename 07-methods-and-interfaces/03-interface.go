package main

import (
	"fmt"
)

// Rectangle is an interface for describing Perimeter method
type Rectangle interface {
	Perimeter() float64
}

// Square is a kind of rectangle
type Square struct {
	length float64
	width  float64
}

// Perimeter method
func (square Square) Perimeter() float64 {
	return square.length*2 + square.width*2
}

// DisplayPerimeter will take a rectangle (so an interface) as parameter
func DisplayPerimeter(rectangle Rectangle) {
	fmt.Println(rectangle.Perimeter())
}

func main() {
	s := Square{length: 2.0, width: 2.0}
	DisplayPerimeter(s)
}
