package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := []Person{
		{Name: "Aman", Age: 22},
		{Name: "Rahul", Age: 25},
		{Name: "Raj", Age: 12},
	}
	fmt.Println("\nbefore change data is :")
	fmt.Println()
	for i, data := range p{
		fmt.Printf("Index %d,Name %s,Age %d\n", i, data.Name, data.Age)
	}

	ptr:=&p[2]
	ptr.Name="Rajesh"
	ptr.Age=30




	fmt.Println("\nafter change data is :")
	fmt.Println()
	for i, data := range p{
		fmt.Printf("Index %d,Name %s,Age %d\n", i, data.Name, data.Age)
	}
}
