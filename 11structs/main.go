package main

import "fmt"

func main() {
	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	yash := User{"Yash", "yash@go.dev", true, 21}

	fmt.Println(yash)
	fmt.Printf("Struct is %+v\n", yash)
	fmt.Printf("Name is %v and Age is %v\n", yash.Name, yash.Age)
}
