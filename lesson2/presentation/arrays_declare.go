package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 23
	a[1] = 44
	fmt.Println(a[0], a[1])
	fmt.Println(&a[0], &a[1])
	fmt.Println(a)
}
