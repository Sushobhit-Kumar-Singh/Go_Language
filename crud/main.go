package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error is getting Response: ", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading : ", err)
	// 	return
	// }
	// fmt.Println("Data : ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding : ", err)
		return
	}
	fmt.Println("Todo: ", todo)
	fmt.Println("Title response : ", todo.Title)
	fmt.Println("completed response : ", todo.Completed)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Shobhit Singh",
		Completed: true,
	}
	// convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling :", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", res)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response ststus : ", res.Status)

}

func performUpdateRequest() {
	todo := Todo{
		UserID:    23789,
		Title:     "Shobhit kumar Golan Hello World",
		Completed: false,
	}
	// convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling :", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	// create PUT Request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT Request : ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response ststus : ", res.Status)

}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	// create Delete Request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request : ", err)
		return
	}

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response ststus : ", res.Status)

}

func main() {
	fmt.Println("Learning CRUD.....")

	performGetRequest()
	performPostRequest()
	performUpdateRequest()
	performDeleteRequest()

}
