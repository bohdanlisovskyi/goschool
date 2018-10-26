package main

import "fmt"

func main() {
	slice := make([]int, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0] = 23
	slice[1] = 55

	fmt.Println(slice)

	//slice[5] = 33 //panic: runtime error: index out of range
}
