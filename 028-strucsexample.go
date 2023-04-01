package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id         int    `json: "id"`
	ParseInt   string `json:"parse_int"`
	GetBoolean bool   `json:"get_boolean"`
}

func main() {

	data := []byte(`{"parse_int":"hola", "get_boolean":true}`)

	var user User
	err := json.Unmarshal(data, &user)

	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(user)
}
