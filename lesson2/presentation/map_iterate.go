package main

import "fmt"

func main() {
	m := map[string]int{
		"first":  1,
		"second": 2,
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}
