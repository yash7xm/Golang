package main

import "fmt"

const LoginToken string = "asdhg" // public

func main() {
	var username string = "Yash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4548455222977
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some alias 
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type 
	var website = "abc.com"
	fmt.Println(website);
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	name := "yash"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}