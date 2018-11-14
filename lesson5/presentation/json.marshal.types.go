package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	i := 12
	m := struct {
		Name   string
		Age    int
		Height float32
		Car    bool
		P      *int
	}{
		Name: "John", Age: 31, Height: 1.72, Car: true, P: &i,
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
