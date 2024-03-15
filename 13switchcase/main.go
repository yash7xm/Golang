package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	diceNum := rand.Intn(6) + 1

	switch diceNum {
	case 1:
		fmt.Println("dice num is 1")
	case 2:
		fmt.Println("dice num is 2")
		fallthrough
	case 3:
		fmt.Println("dice num is 3")
	default:
		fmt.Println("this is just default")
	}
}
