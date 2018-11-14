package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	m := User{
		Name: "John",
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
