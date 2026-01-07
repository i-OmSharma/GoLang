package main

import (
	"fmt"
	"time"
)

func main() {
	
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)
	
	
	formatted := currentTime.Format("02-01-2006, Monday")
	fmt.Println("Formatted Time: ", formatted)
	
	layout_str := "2006-01-02"
	timeStr := "2026-01-07"
	formatted_time, _ := time.Parse(layout_str, timeStr)
	fmt.Println("Formatted tiem converted: ", formatted_time)
	
	//add 1 more day!
	// 
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new date is: ", new_date)
	formatted_date := new_date.Format("02-01-2006")
	fmt.Println("Formatted date: ", formatted_date)
	
	//reduce 1 day
	prev_date := currentTime.Add(-24 * time.Hour)
	fmt.Println("new date is: ", prev_date)
	formatted_date2 := prev_date.Format("02-01-2006")
	fmt.Println("Formatted date: ", formatted_date2)
}