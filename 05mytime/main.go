package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to go lang time study")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2024, time.June, 23, 15, 05, 0, 0, time.UTC)
	fmt.Println(createdTime.Format("2006-01-02"))
}
