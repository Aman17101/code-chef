package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Direct interface call without storing in a variable
	fmt.Println("Area of circle is:", Shape(Circle{radius: 5.0}).Area())
}
