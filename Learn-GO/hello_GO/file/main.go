package main

import (
	"fmt"
	"os"
)

func main() {

	/* File Creation

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

	*/

	/* Read file content -> manual
	file, err := os.Open("example.txt")
	if err != nil{
		fmt.Println("Error while opening file", err)
		return
	}
	defer file.Close()

	// Create a buffer to read File content

	buffer := make([]byte, 1024)

	// read the file content ino the buffer -> it returns integer or error

	for {
		n, err := file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("error while reading: ", err)
			return
		}

		// reading the file.
		fmt.Println(string(buffer[:n])) // read till 0 to n -> n here integer
	}
	*/

	// content, err := ioutil.ReadFile("example.txt") // ioutil is deprecated
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	fmt.Println(string(content))
}
