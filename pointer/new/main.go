package main

import "fmt"

func main(){
	p:=new(int )

	*p=11

	fmt.Println("value of pointer is :",*p)

}