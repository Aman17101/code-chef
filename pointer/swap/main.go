package main

import "fmt"

func swap(a, b *int) {
	// your code here
	*a, *b = *b, *a

}

func main() {
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y) // Output: x: 20 y: 10
}
