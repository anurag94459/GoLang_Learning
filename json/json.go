package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println()
	person1 := person{Name: "Anurag", Age: 32}
	fmt.Println(person1)

	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error in marshaling")
	}

	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	var person2 person
	error := json.Unmarshal(jsonData, &person2)
	if error != nil {
		fmt.Println("Error Unmarshaling")
		return
	}
	fmt.Println(person2)
}
