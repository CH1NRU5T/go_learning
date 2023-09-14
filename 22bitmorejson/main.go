package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"` // json will have key <courseName> not <Name>
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// json encoding
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LCO", "password", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LCO", "password123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 399, "LCO", "password321", nil},
	}
	// package this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
