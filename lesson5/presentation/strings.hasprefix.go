package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Maytech", "May"))
	fmt.Println(strings.HasPrefix("Maytech", "may"))
	fmt.Println(strings.HasPrefix("Maytech", ""))
}
