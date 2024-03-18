package main

import (
	"fmt"
	"net/url"
)

const myUrl = `https://lco.dev:3000/learn?course="React"`

func main() {
	fmt.Println("Welocme to handling url")

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
}
