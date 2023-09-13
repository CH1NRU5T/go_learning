package main

import "fmt"
const LoginToken string = "e;srjnfw;or";
func main(){
	var isLoggedIn bool = true;
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type: %T \n", isLoggedIn);
	var smallValue int = 255;
	fmt.Println(smallValue);
	fmt.Printf("Variable is of type: %T \n", smallValue);
	var smallFloat float64 = 255.01265165165165115514;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n", smallFloat);
	
	// default values and aliases
	var anotherVariable string ;
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type: %T \n", anotherVariable);

	// no var style
	numberOfUser := 30000;
	fmt.Println(numberOfUser);
	

	// implicit way of declaring variables
	var website = "www.google.com";
	fmt.Println("Website is: ", website);

	fmt.Println(LoginToken)
}