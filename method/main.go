package method

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	area := math.Pi * c.radius * c.radius
	return area

}
func main() {
	c:=Circle{radius :6.99}
	fmt.Println("Area of circle is :",c.Area())
}