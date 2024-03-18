package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welocme to files")

	file, err := os.Create("./myfile.txt")

	// if err != nil  {
	// 	panic(err)
	// }
	checkNilError(err)

	content := "Writing in the new made file"

	length, err := io.WriteString(file, content)

	fmt.Println("Length of the new file is: ", length)
	defer file.Close()
	readFile("myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("New file content is ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
