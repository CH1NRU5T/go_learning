package main

import "fmt"

func main() {
	fmt.Println("Maps")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("JS shorts for", languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
