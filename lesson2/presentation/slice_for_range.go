package main

import "fmt"

func main() {
	friends := []int{77, 88, 99, 66, 55, 44, 33, 99}
	for i, v := range friends {
		fmt.Printf("%d : v[%d]\n", i, v)
	}
}
