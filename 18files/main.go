package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	content := "This needs to go in a file - lco.in"
	file, err := os.Create("./log.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	// error := os.Remove("./log.txt")
	// if error != nil {
	// 	panic(error)
	// }
	file.Close()
	// defer os.Remove("./log.txt")
	readFile("./log.txt")
}

func readFile(fileName string) {
	db, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(db))
}
