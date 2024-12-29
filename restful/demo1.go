package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	//jsonTodo, _ := json.Marshal(todo)
	jsonTodo, _ := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
