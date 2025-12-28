package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main(){
	
	// reader := bufio.NewReader(os.Stdin)
	// num := reader.ReadSlice('\n')
	num := []int{1,2,3,4,5,6,7,8,9}
	
	fmt.Println("slice num = ", num)
	fmt.Printf("Data Type = %T\n", num)
	fmt.Println("Length", len(num))
	
	num = append(num, 23,34,45)
	fmt.Println("slice num = ", num)
	
	all := make([]int, 3, 5)
	fmt.Println("slice all = ", all)
	fmt.Printf("Data Type = %T\n", all)
	fmt.Println("Length", len(all))
	
}