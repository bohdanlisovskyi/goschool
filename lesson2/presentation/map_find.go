package main

import "fmt"

func main() {
	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	v, ok := m["fourth"]
	fmt.Println("fourth", ok, v)

	v, ok = m["second"]
	fmt.Println("second", ok, v)
}
