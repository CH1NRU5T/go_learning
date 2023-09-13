package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?courseName=reactjs&paymentId=alsb5doi12ub"

func main() {
	// fmt.Println("Urls")
	// fmt.Println(myUrl)
	_, error := url.Parse(myUrl)
	if error != nil {
		panic(error)
	}
	// fmt.Println(url.Scheme)
	// fmt.Println(url.Host)
	// fmt.Println(url.Path)
	// fmt.Println(url.Port())
	// fmt.Println(url.RawQuery)

	// qparams1

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=ansh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
