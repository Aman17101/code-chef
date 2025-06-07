package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	max := arr[0]
	for i := 0; i <= len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println( "maximum value is :",max)
}
