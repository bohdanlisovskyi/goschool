package main

import "fmt"

func main() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	fmt.Println(foo)
}
