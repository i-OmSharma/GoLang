package main

import "fmt"

type Person struct {
	FristName string
	LastName  string
	Age       int
}

func main() {
	var Om Person

	Om.FristName = "Om"
	Om.LastName = "SHA"
	Om.Age = 21

	fmt.Println("Person Om", Om)
	
	// Second Method
	Person1 := Person {
		FristName: "Akash",
		LastName: "Bikash",
		Age: 23,
	}
	
	fmt.Println("Person1 details : ", Person1)
	
	
}
