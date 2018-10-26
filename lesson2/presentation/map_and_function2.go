package main

import "fmt"

func main() {
	m := make([]int, 1, 6)
	fmt.Println(m)

	AddKeyValue(m)

	fmt.Println(m)
}

func AddKeyValue(m []int) {
	m = append(m, 1)
}
