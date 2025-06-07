package main

import "fmt"

func main() {
	integer := []int{1,2,3,4,5}
	fmt.Println("before change")
	for i, val := range integer {
		fmt.Println("integer index is :",i,"integer val is" ,val)
	}
 ptr :=&integer[2]


 *ptr=33


 fmt.Println("after change")
 for i, val := range integer {
		fmt.Println("integer index is :",i,"integer val is" ,val)

 }
	
	
	

}

