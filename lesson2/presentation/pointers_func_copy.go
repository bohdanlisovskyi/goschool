package main

import "fmt"

func main() {

	i := 1
	incr(i)
	fmt.Println(i)
}

func incr(i int) {
	i++
}
