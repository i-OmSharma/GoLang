package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web requests")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error gtting GET resonse", err)
		return
	}
	defer res.Body.Close()
	
	fmt.Printf("Type of response: %T\n", res)
	// fmt.Println("Response: ", res)
	
	// Read the response
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error erading response body: ", err)
		return
	}
	fmt.Printf("body type: %T\n",data)
	fmt.Println("response: ",string(data))


}
