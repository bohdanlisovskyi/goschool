package main

import "fmt"

func main() {
	slice := []int{3, 4, 5, 6, 7, 8}

	fmt.Println(len(slice))

	sl := slice[2:5]

	fmt.Println(sl)

	fmt.Println(len(sl))
}
