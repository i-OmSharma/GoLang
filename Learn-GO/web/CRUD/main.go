package main

import "fmt"

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func main() {
	fmt.Println("CRUD Operation")
	PerformGET()
	// PerformPOST()
	// PerformUPDATE()
	// PerformDELETE()
}
