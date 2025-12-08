package main

import "fmt"

func main() {
	fmt.Println("Start of error handling function ")
	
//--------------with err handling-------------------------//
	// ans, err := divide(10,3)
	// if err != nil{
	// 	fmt.Println(" got an error")
	// }
//-------------wihout err, use of (_)--------------------//
	ans, _ := divide(10,3)
	fmt.Println(ans)
}

func divide(a, b float64) (float64, error) {
	if b ==0 {
		return 0, fmt.Errorf("Error handling")
	}; return a/b, nil
}

func divide2(a, b float64) (float64, string) {
	if b ==0 {
		return 0, "Error handling"
	}; return a/b, "nil"
}