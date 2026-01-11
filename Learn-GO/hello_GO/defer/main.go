package main

import "fmt"

func add(a,b int) int {
	return a + b
}

func main() {
	fmt.Println("start of program")
	data := add(5, 6)
	fmt.Println("Data is: ", data)
	defer fmt.Println("Middle of program")
	fmt.Println("end of program")
}
