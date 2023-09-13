package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in GoLang")
	result, _, _ := proAdder(3, 5, 1, 1, 1)

	fmt.Println("Result is: ", result)
}

func proAdder(values ...int) (int, string, bool) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "HI", true
}

func adder(one int, two int) int {
	return one + two
}

func greeter() {
	fmt.Println("Namastey Go")
}
