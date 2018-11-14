package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Maytech", "ech"))
	fmt.Println(strings.HasSuffix("Maytech", "ECH"))
	fmt.Println(strings.HasSuffix("Maytech", ""))
}
