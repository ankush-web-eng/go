package main

import (
	"encoding/json"
	"fmt"
)

func throwErr(err error) {
	if err != nil {
		panic(err)
	}
}

type name struct {
	Name   string `json:"student"`
	Age    int
	Course string
	Branch string
}

func main() {
	fmt.Println("This is JSON Package")
	// encodeJson()
	checkJson()
}

func encodeJson() {
	ankush := []name{
		{"Ankush", 19, "B.Tech", "CSE"},
		{"Abhay", 19, "B.Tech", "CSE"},
		{"Abhishek", 19, "B.Tech", "CSE"},
		{"Harshit", 19, "B.Tech", "CSE"},
	}

	data, err := json.MarshalIndent(ankush, "", "\t")
	throwErr(err)

	fmt.Printf("%s\n", data)
}

func checkJson() {
	data := []byte(`
	{
		"student": "Ankush",
		"Age": 19,
		"Course": "B.Tech"
	}
	`)

	var student name

	isValid := json.Valid(data)
	if isValid {
		json.Unmarshal(data, &student)
		fmt.Printf("%#v\n", student)
	} else {
		fmt.Println("Invalid JSON")
	}
}
