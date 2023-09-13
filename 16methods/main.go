package main

import "fmt"

func main() {
	fmt.Println("Methods")
	ansh := User{"Ansh", "ansh@gmail.com", true, 15}
	ansh.GetStatus()
	ansh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
