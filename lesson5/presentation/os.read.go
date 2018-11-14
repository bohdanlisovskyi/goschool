package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("/Users/bohdan/projects/gomaytech/src/golangschool/lesson5/presentation/test.txt")

	b1 := make([]byte, 5)
	n1, _ := f.Read(b1)

	fmt.Printf("%d bytes: %s\n", n1, b1)
}
