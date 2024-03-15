package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i := range arr {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println("")

	for i,val := range arr {
		fmt.Printf("Index is %v and value is %v\n", i, val)
	}

	for _, val := range arr {
		if val > 10 {
			goto bigNum
		}
	}

	bigNum:
		fmt.Println("this is big num")
}
