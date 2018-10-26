package main

import "fmt"

func main() {
	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}
	fmt.Println(m)

	AddKeyValue(m)

	fmt.Println(m)
}

func AddKeyValue(m map[string]int) {
	m["asdad"] = 123
}
