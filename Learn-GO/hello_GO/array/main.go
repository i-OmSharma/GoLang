package main

import "fmt"

func main() {
	fmt.Println("Arrays in GO")

	var name [5]string

	name[0] = "Alen"
	name[1] = "liza"

	fmt.Println("Name of Persons are:", name)
	
	
	var num = [5]int{1,2,3,4}
	fmt.Println("Count =", num)
	
	arr := []int{6,7,8,9}
	fmt.Println("arr = ", arr)
	
	fmt.Println("length of aarray =", len(arr))
	
	fmt.Println("Arrray at index 2 = ", arr[2])
}
