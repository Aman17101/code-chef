package main

import "fmt"

func doubleValue(x *int) {
	// your code here

	* x=2*(*x)
}

func main() {
	num := 5
	doubleValue(&num)
	fmt.Println(num) // Output: 10
}
