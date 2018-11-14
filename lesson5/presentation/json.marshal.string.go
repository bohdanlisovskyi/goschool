package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
}

func main() {
	m := User{
		Name: "<John & Luis>",
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
