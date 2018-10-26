package main

import "fmt"

func main() {
	a := [3]int{33, 44, 55}

	for i, v := range a {
		fmt.Println(i, v)
	}
}
