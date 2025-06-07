package main

import "fmt"

func main() {
	value := []int{}

	value = append(value, 1)
	value = append(value, 2)
	value = append(value, 3)
	value = append(value, 4)
	value = append(value, 5)
	value = append(value, 6)
	value = append(value, 7)
	value = append(value, 9)
	value = append(value, 8)

	if len(value) == 0 {
		fmt.Println("Slice is empty")
	}

	max := value[0]
	index := 0

	for i, val := range value {
		fmt.Println(val)
		if val > max {
			max = val
			index = i
		}

	}
	fmt.Println("index of maximum value is :", index)
	fmt.Println("maximum value is :", max)
}
