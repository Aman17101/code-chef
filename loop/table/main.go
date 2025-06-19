package main

import "fmt"

func main() {

	table := [10]int{}

	var  x int

	fmt.Print("Enter a number to print its multiplication table: ")
	fmt.Scanln(&x)      // take user input and store it in x

	for i := 0; i < 10; i++ {
		table[i] = x * (i + 1)
	}

	for i, j := range table {
		fmt.Println("table of" , x, "is:",x,"*", i+1,"=" ,j)
	}

}
