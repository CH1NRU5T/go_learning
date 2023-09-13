package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Something else"
	}
	fmt.Println(result)
	if num := 3; num < 10 {
		fmt.Println("< 10")
	} else {
		fmt.Println(">= 10")
	}

}
