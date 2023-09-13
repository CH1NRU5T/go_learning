package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case example")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("2 spot")
	case 3:
		fmt.Println("3 spot")
		fallthrough
	case 4:
		fmt.Println("4 spot")
	case 5:
		fmt.Println("5 spot")
	case 6:
		fmt.Println("6 spot, go again")
	}
}
