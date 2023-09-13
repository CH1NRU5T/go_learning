package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent
	ansh := User{"Ansh", "chinrust@gmail.com", true, 21}
	fmt.Printf("User: %+v\n", ansh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
