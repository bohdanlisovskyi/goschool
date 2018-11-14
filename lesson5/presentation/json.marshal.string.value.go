package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int `json:"age,string,omitempty"`
}

func main() {
	m := User{
		Name: "John",
		Age:  31,
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
