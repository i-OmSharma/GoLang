package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
)

type Todo struct {
	Id int `json:"id"`
	Todo string `json:"todo"`
	Completed bool `josn:"completed"`
	UserId int `json:"userId"`
}

func main() {
	fmt.Println("CRUD Operation")
	
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
	
	err = json.Unmarshal(&todo)
	if err != nil{
		fmt.Println("error while Unmarshellig: ", err)
		return
	} 
	fmt.Println("Decoded data is: ",)
	
}