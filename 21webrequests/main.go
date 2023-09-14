package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// performGetRequest("http://localhost:8000/get")
	// performPostRequest("http://localhost:8000/post")
	performPostFormRequest("http://localhost:8000/postform")

}
func performGetRequest(myUrl string) {
	// const myUrl = "http://localhost:8000/get";
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println(byteCount)
}
func performPostFormRequest(myUrl string) {
	// formdata
	data := url.Values{}
	data.Add("firstName", "Ansh")
	data.Add("lastName", "Agrawal")
	data.Add("email", "ansh@gmail.com")

	response, err := http.PostForm(myUrl, data)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
func performPostRequest(myUrl string) {
	// const myUrl = "http://localhost:8000/get";
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"courseName":"LCO GoLang",
			"price":12,
			"platform":"google.com"
			
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println(byteCount)
}
