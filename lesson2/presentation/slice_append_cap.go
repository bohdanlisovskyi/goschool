package main

import "fmt"

func main() {
	a := make([]int, 12)
	fmt.Println("a:", a)
	fmt.Println("addr a[0]: ", &a[0])
	fmt.Println("a cap: ", cap(a))
	a = append(a, 1)
	fmt.Println("new a: ", a)
	fmt.Println("new addr a[0]", &a[0])
	fmt.Println("new a cap: ", cap(a))
}
