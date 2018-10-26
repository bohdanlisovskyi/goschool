package main

import "fmt"

func main() {
	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	fmt.Println(m)

	delete(m, "second")

	fmt.Println(m)
}
