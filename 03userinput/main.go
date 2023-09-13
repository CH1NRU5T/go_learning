package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	welcome := "Welcome to user input";
	fmt.Println(welcome);
	reader := bufio.NewReader(os.Stdin);
	fmt.Print("Enter the rating for our pizza: ");

	// comma ok, error err
	input, _ := reader.ReadString('\n');
	fmt.Println(input);
	fmt.Printf("Type of rating: %T \n", input);

}