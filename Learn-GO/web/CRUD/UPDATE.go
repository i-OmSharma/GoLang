package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func PerformUPDATE() {

	todo := Todo{
		Id:        21,
		Todo:      "Play game",
		Completed: true,
		UserId:    12,
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

	//Create PUT request!
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error while creatign PUT request: ", err)
		return
	}
	req.Header.Set("content-type", "applicaiton/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response code: ", res.StatusCode)

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
}
