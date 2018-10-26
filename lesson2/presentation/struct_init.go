package main

import "fmt"

type House struct {
	Windows int
	Doors   int32
	Name    string
}

func main() {
	h := House{
		Windows: 12,
		Doors:   1,
		Name:    "best house",
	}
	fmt.Printf("%+v\n", h)
}
