package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`{"name": "Platypus", "age": 12}`)
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var u User
	err := json.Unmarshal(jsonBlob, &u)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%T %+v", u, u)
}
