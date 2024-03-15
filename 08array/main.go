package main

import "fmt"

func main() {
	fmt.Println("Array comes here")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
}
