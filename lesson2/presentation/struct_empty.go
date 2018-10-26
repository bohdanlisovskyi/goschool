package main

import (
	"fmt"
	"unsafe"
)

type Empty struct{}

func main() {

	a := Empty{}
	fmt.Printf("%+v\n", unsafe.Sizeof(a))

	i := true
	fmt.Printf("%+v\n", unsafe.Sizeof(i))
}
