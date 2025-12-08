package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Whats ur name???")

	//--------Basic Method!!!---------------//
	// var name string
	// fmt.Scan(&name)
	// fmt.Print("Hello,", name)
	//-------------------------------------//

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Print("Hello,", name)
}
