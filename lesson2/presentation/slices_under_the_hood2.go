package main

import "fmt"

func main() {
	foo := make([]int, 5)
	fmt.Println("foo: ", foo)

	foo[3] = 42
	foo[4] = 100

	fmt.Println("foo: ", foo)

	bar := foo[1:4]
	fmt.Println("bar: ", bar)

	bar[1] = 99
	fmt.Println("bar: ", bar)
	fmt.Println("foo: ", foo)
}
