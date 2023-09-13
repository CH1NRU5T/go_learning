package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome tovideo on slices")
	// var fruits = []string{"Apple", "Tomato", "Peach", "Banana", "Grapes"}
	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777
	highScore = append(highScore, 555, 666, 321)

	// fmt.Println(highScore)
	// sort.Ints(highScore)
	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	// remove value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
