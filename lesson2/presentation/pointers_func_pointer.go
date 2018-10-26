package main

import "fmt"

func main() {

	i := 1
	incrPointer(&i)
	fmt.Println(i)
}

func incrPointer(i *int) {
	*i++
}
