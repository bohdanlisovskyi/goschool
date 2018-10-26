package main

import "fmt"

func main() {
	h := struct {
		windows int
		doors   int
	}{
		windows: 123,
		doors:   323,
	}

	fmt.Printf("%+v\n", h)
}
