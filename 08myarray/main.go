package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Tomato"
	fruits[3] = "Peach"
	fmt.Println(len(fruits))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

}
