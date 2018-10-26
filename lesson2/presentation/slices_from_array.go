package main

import "fmt"

func main() {
	ar := [3]int{3, 4, 5}
	fmt.Println(ar)
	sl := ar[:]
	fmt.Println(sl)
}
