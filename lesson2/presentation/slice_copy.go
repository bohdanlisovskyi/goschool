package main

import "fmt"

func main() {

	t := make([]int, 4)
	s := []int{1, 2, 2, 3, 5, 512}

	copy(t, s)

	fmt.Println(&t[0], &s[0])
	fmt.Println(t)
}
