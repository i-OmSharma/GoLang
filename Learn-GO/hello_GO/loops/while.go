package main

import "fmt"

func while() {
	i := 0
	
	for {
		fmt.Println("While loop")
		i++
		if i == 3 {
			break
		}
	}
}