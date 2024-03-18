package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welocme to files")

	file, err := os.Create("./myfile.txt")

	if err != nil  {
		panic(err)
	}

	content := "Writing in the new made file"

	length, err := io.WriteString(file, content)

	fmt.Println("Length of the new file is: ", length)

}