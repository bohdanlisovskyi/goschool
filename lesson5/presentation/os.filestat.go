package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	st, _ := f.Stat()
	fmt.Printf("%+v\n", st)
}
