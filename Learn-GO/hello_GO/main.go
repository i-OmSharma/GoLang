package main

import (
	"fmt"
	"hello_GO/myutil"
)

func main() {
	fmt.Println("Hello GO...")

	myutil.PrintMessage("hi GO...")

	var name string = "Om"
	var version = 23

	fmt.Println(name + fmt.Sprint(version))
}
