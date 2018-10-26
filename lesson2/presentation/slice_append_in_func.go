package main

import "fmt"

func main() {
	n := []int{7, 6, 5, 4, 3, 2}
	fmt.Println(n)
	Append(n)
	fmt.Println(n)
}

func Append(n []int) {
	n = append(n, 2)
}
