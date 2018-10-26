package main

import "fmt"

func main() {
	var h struct {
		windows int
		doors   int
	}

	fmt.Printf("%+v\n", h)
}
