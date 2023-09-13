package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Hello World")
	defer fmt.Println("Hello World 1")
	fmt.Println("Wow")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
