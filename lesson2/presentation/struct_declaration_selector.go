package main

import "fmt"

type House struct {
	Windows int
	Doors   int32
	Name    string
}

func main() {
	h := House{
		Name:    "best house",
		Doors:   2123,
		Windows: 133,
	}

	fmt.Println("Doors: ", h.Doors)
	fmt.Println("Windows: ", h.Windows)
	fmt.Println("Name: ", h.Name)
}
