package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Functions in GoLang \n")
	simpleFunction()
	
	sum := add()
	fmt.Println("Sum is ", sum)
	
	mul := multiply()
	fmt.Println("multiplication is ", mul)
}  


func simpleFunction() {
	fmt.Println("Simple function")
}

func add() (result int) {
	
	fmt.Println("Enter a")
	var a int
	fmt.Scan(&a)
	fmt.Println("Enter b")
	var b int 
	fmt.Scan(&b)
	
	result = a + b
	return 
}

func multiply() (result2 int) {
	fmt.Println("Enter a")
	var a int
	fmt.Scan(&a)
	fmt.Println("Enter b")
	var b int 
	fmt.Scan(&b)
	
	result2 = a * b
	return
}