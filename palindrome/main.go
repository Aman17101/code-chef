package main

import "fmt"

func isPalindrome(x int) bool {
	cons := 0
	temp := x

	for temp > 0 {
		digit := temp % 10
		cons = 10*cons + digit
		temp = temp / 10

	}

	if cons == x {
		return true
	} else {
		return false

	}

}

func main() {
x :=123

	res:=isPalindrome(x)
	fmt.Println("palindrome is",res)
}