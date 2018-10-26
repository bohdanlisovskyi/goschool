package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["first"] = 1
	m["second"] = 2

	fmt.Println(m)
}
