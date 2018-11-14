package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndex("MaytechMaytech", "a"))
	fmt.Println(strings.LastIndex("MaytechMaytech", "ayt"))
	fmt.Println(strings.LastIndex("MaytechMaytech", "z"))
}
