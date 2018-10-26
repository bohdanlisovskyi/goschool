package main

import "fmt"

func main() {
	m := map[string]int{
		"first":  1,
		"second": 2,
	}

	for key := range m {
		fmt.Println(key)
	}
}
