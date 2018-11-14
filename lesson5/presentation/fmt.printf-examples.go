package main

import "fmt"

func main() {
	var b bool = true
	var i int = 124
	var str string = "some text"
	var p = &str
	var t = struct {
		Value string
	}{
		Value: str,
	}

	fmt.Printf("%t %d %s %p\n", b, i, str, p)
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%T", t)
}
