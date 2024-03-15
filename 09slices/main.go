package main

import "fmt"

func main() {
	var list = []string{"a", "b", "c"}

	fmt.Println("List before appending: ", list)
	list = append(list, "d", "e")
	fmt.Println("List after appending: ", list)

	list = (list[1:3])
	fmt.Println("List after slice using colon", list)

	// using make keyword to create slices
	highScores := make([]int, 2)

	highScores[0] = 100
	highScores[1] = 99

	// append will reallocate the memory so we are able to append
	// even if the size we indicated is only 2
	highScores = append(highScores, 80, 95)

	// now length is also increased
	fmt.Println("Highscores: ", highScores, len(highScores))

	// removing an element using slice
	var courses = []string{"cpp", "java", "go", "swift", "python"}
	index := 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after deleting: ", courses)
}
