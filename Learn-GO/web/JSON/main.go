package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON data")
	person1 := Person{Name: "Jhon", Age: 21, IsAdult: true}
	fmt.Println("Person no 1 details: ", person1)
	
	// Converting into JSON (marsheling)
	jsonData, err := json.Marshal(person1) // returns slice of bytes and error
	if err != nil {
		fmt.Println("Error during Marshelling", err)
		return
	}
	fmt.Println("Person data is: ", string(jsonData))
	
	//Decoding
	var PersonData Person
	err = json.Unmarshal(jsonData, &PersonData)
	if err != nil{
		fmt.Println("error while Unmarshellig: ", err)
		return
	}
	fmt.Println("Decoded data is: ", PersonData)
}

// Normal struct for
// type Person struct {
// 	Name string 
// 	Age int
// 	isAdult bool
// }

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}