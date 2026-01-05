package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 12
	fmt.Println("Number is: ", num)
	fmt.Printf("Type of num is %T\n", num)
	
	// num = num + 1.12
	
	var data float64 = float64(num)
	data = data + 1.12
	fmt.Println("Number is: ", data)
	fmt.Printf("Type of num is %T\n", data)
	
	
	num1 := 123
	str := strconv.Itoa(num1)
	fmt.Println("str is: ", str)
	fmt.Printf("Type of str is %T\n", str)
	
	str1 := "1234"
	num2, _ := strconv.Atoi(str1)
	fmt.Println("num2 is: ", num2)
	fmt.Printf("Type of num2 is %T\n", num2)
	
	str2 := "1234.4"
	num3, _ := strconv.ParseFloat(str2, 64)
	fmt.Println("num3 is: ", num3)
	fmt.Printf("Type of num3 is %T\n", num3)
	
	
	
}