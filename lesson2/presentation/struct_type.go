package main

import (
	"fmt"
)

type NameOfStruct struct {
	i int
	f float32
	s string
}

func main() {

	s := NameOfStruct{}
	fmt.Printf("output : %+v\n", s)
}
