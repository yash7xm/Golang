package main

import "fmt"

func main() {
	languages := make(map[string]string)

	// creating key value pairs
	languages["JS"] = "Javascript"
	languages["RB"] = "Rubby"
	languages["PY"] = "Python"

	fmt.Println("All the languages are: ", languages)
	fmt.Println("JS is short for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("After deleting Rubby", languages)

	// looping through the map

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
