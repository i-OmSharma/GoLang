package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 5")
	} else if z > 5 {
		fmt.Println("z is less than 10  but more then or elquals to 5")
	} else {
		fmt.Println("z is less then 5")
	}

	y := 10
	if x > 5 && (y > 5 || z < 10) {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}
	
	// y := 10
	// if (x > 5) && ((y > 5 || z < 10)) {
	// 	fmt.Println("x is greater than 5")
	// } else {
	// 	fmt.Println("x is less than 5")
	// }
}
