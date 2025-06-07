package main

import "fmt"

type Person struct {
	name string
	age  int
}

func updatePerson(p *Person) {
	// Change name to "Updated" and age to 30

	p.name="Aditya"
	p.age = 30
}

func main() {
	p := Person{name: "Aman", age: 22}
	updatePerson(&p)
	fmt.Println(p) // Output: {Updated 30}
}
