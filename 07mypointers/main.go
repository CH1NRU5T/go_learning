package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("The value of pointer is:", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is:", ptr)
	fmt.Println("Value of actual pointer is:", &ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value:", myNumber)
}
