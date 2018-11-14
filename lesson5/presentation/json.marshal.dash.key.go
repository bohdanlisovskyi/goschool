package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"-,"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	m := User{
		Name: "John",
		Age:  123,
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
