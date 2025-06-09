package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("not divisible")
	}
	return a / b, nil

}
func main() {

	a, b := 10, 11
	result, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("result", result)

}
