package main

import "fmt"

type Street struct {
	Windows int
	House1
}

type House1 struct {
	Windows int
	Doors   int
}

func main() {
	h := Street{}
	fmt.Printf("%+v\n", h)

	h.Windows = 2
	h.Doors = 1

	fmt.Printf("%+v\n", h)
}
