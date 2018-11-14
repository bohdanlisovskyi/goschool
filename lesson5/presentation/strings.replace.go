package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Maytech is Maytech", "ch", "cher", -1))
	fmt.Println(strings.Replace("Maytech is Maytech", "ch", "cher", 0))
	fmt.Println(strings.Replace("Maytech is Maytech", "ch", "cher", 1))
}
