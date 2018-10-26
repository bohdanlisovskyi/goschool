package main

import "fmt"

func main() {
	var d []int

	fmt.Println("cap: ", cap(d))
	fmt.Println("len: ", len(d))

	d = append(d, 1)

	fmt.Println("cap: ", cap(d))
	fmt.Println("len: ", len(d))
	fmt.Println(d)
}
