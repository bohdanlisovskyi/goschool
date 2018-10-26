package main

import "fmt"

type House struct {
	Windows int
}

func main() {
	h := House{
		Windows: 133,
	}

	fmt.Println("Windows: ", h.Windows)

	h.Windows = 666

	fmt.Println("Windows: ", h.Windows)
}
