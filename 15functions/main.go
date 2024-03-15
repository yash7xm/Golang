package main

import "fmt"

func main() {
	result := add(2, 3)
	fmt.Println(result)

	proRes := proAdder(1, 2, 3, 4, 5)
	fmt.Println(proRes)

	sum, msg, sum2 := sumMsg()
	fmt.Println(sum, msg, sum2)
	
}

func sumMsg() (int, string, int) {
	return 1, "abc", 2
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}

func add(a int, b int) int {
	return a + b
}
