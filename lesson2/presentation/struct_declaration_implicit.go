package main

import "fmt"

type House struct {
	Windows int
	Doors   int32
	Name    string
}

func main() {
	h := House{
		Name: "best house",
	}
	fmt.Printf("%+v\n", h)
}
