package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	// "io"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func PerformGET() {
	res, err := http.Get("https://dummyjson.com/todos/1")
	if err != nil {
		fmt.Println("Error while fetching data: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error while getting Response: ", res.Status)
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error while reading data: ", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	// var result map[string] interface{}

	// err = json.Unmarshal(data, &result)
	// if err != nil {
	// 	fmt.Println("error while Unmarshalling: ", err)
	// 	return
	// }
	// fmt.Println("decoded data is: ", result)

	var todo Todo

	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error while decoding: ", err)
	}
	fmt.Println("Todo is : ", todo)
}

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

func main() {
	fmt.Println("CRUD Operation")
	// PerformGET()
	PerformPOST()

}
