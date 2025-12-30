package main

import (
	"fmt"
	"strings"
)
		

func range_loop() {
	// num := []int{31,42,73,14,95,86}

	// for index, value := range num {
	// 	fmt.Printf("index of number is %d, and value is %d \n", index, value)
	// }

	
	//------------------with each character----------------------------//
	// name := "raj, mass, bikash"

	// for index, char := range name {
	// 	fmt.Printf("index of name is %d and value is %c \n", index, char)
	// }
	
	//-----------------------wit seperate character----------------------//
	
	name := "raj, mass, bikash"

	names := strings.Split(name, ", ")

	for index, char := range names {
		fmt.Printf("index of name is %d and value is %s \n", index, char)
	}
}
