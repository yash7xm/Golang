package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Http request module")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response we get: %T\n", response)
	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(databyte))

}
