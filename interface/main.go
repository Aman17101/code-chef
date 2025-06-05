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
	area := math.Pi *c.radius*c. radius
	return area
}

func main() {
	c := Circle{radius: 5.0}
var s Shape=c
//	s := c
	fmt.Println("Area of circle is :", s.Area())
}
