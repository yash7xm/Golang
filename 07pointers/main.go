package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("Pointer is : ", ptr)

	num := 23
	// var ptr *int = &num
	ptr := &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(*ptr)
}
