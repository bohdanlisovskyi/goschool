package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	m := User{
		Name: "John",
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
