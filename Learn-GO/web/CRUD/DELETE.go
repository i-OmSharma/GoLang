package main

import (
	"fmt"
	"net/http"
)

func PerformDELETE() {
	myURL := "https://dummyjson.com/todos/1"

	//create DELETE request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error while creatign DELETE request: ", err)
		return
	}
	//Create client
	client := http.Client{}
	//send request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response code: ", res.StatusCode)

	// data, _ := io.ReadAll(res.Body)
	// fmt.Println("Response: ", string(data))

}
