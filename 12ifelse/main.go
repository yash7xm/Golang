package main

import "fmt"

func main() {
	value := 10

	if value > 10 {
		fmt.Println("Value is greater than 10")
	} else if value < 10 {
		fmt.Println("value is less than 10")
	} else {
		fmt.Println("value is 10")
	}

	if a := 5; a < 10 {
		fmt.Println("a is less than 10")
	} else {
		fmt.Println("a is not less than 10")
	}
}
