package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Printf("Response is of type %T\n", response)
	// databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// content := string(databytes)
	fmt.Println(response.Status)
}
