package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil{
		fmt.Println("Error while Creating file", err)
		return
	}
	defer file.Close()
	
	content := "this file management class"
	// _, errors := io.WriteString(file, content+"\n") //+\n is for move cursor to next line.
	byte, err := io.WriteString(file, content)
	
	fmt.Println("Bytes written: ", byte)
	if err != nil{
		fmt.Println("Error while Creating file", err)
		return
	}
	
	
	
	fmt.Println("Successfully created a file")
}