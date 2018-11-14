package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	name string
}

func main() {
	m := User{
		name: "John",
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
