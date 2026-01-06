package main

import (
	"fmt"
	"strings"
)

func main() {
	fruit := "apple, orange, banana"
	parts := strings.Split(fruit, ",")
	fmt.Println("specific fruit: ",parts)
	
	
	num := "one,two, three, four, two, five"
	count := strings.Count(num, "two")
	fmt.Println("Count of two is: ", count)
	
	
	str := "               HELLO,                  GO!                        "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed string is: ", trimmed)
	
	str = "               HELLO,                  GO!                        "
	joined := strings.Join(strings.Fields(str), " ")
	fmt.Println("Joined string is: ", joined)
	
	
	str1 := "Om"
	str2 := "SHA"
	result := strings.Join([]string{str1, str2}," ")
	fmt.Println("Final name is: ", result)
	
}