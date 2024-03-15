package main

import "fmt"

type User struct {
    Name   string
    Email  string
    Status bool
    Age    int
}

func main() {
    // Creating an instance of User
    user := User{
        Name:   "John",
        Email:  "john@example.com",
        Status: true,
        Age:    30,
    }

    // Calling GetStatus method on the user instance
    user.GetStatus()
}

func (u User) GetStatus() {
    fmt.Println("Is user active: ", u.Status)
}
