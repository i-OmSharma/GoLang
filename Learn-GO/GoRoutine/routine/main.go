package main

import (
	"fmt"
	"time"
)


func sayHello(){
	fmt.Println("Hello from Go!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello agaain from Go!")
}

func sayHi() {
	fmt.Println("Hi from Go!")
}


func main() {
	fmt.Println("GoRoutine")
	go sayHello()
	sayHi()
	
	//Let goRoutine work!
	time.Sleep(1000 * time.Millisecond)
}