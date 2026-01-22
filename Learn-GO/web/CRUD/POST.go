package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PerformPOST() {
	todo := Todo{
		Id:        23,
		Todo:      "Play game",
		Completed: false,
		UserId:    21,
	}

	//Convert it into JSON--> Marshel

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error during Marshelling", err)
		return
	}
	fmt.Println("Person data is: ", string(jsonData))
	//convert Data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	//Send POST reaquest
	myURL := "https://dummyjson.com/todos"
	res, err := http.Post(myURL, "application/json", jsonReader) //urlstring, content type, reader
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response code: ", res.StatusCode)

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))

}
