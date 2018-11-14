package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("A B : ", strings.Compare("A", "B"))                         // A < B
	fmt.Println("B A : ", strings.Compare("B", "A"))                         // B > A
	fmt.Println("Japan Australia : ", strings.Compare("Japan", "Australia")) // J > A
	fmt.Println("Australia Japan : ", strings.Compare("Australia", "Japan")) // A < J
	fmt.Println("Germany Germany : ", strings.Compare("Germany", "Germany")) // G == G
	fmt.Println("Germany GERMANY : ", strings.Compare("Germany", "GERMANY")) // GERMANY > Germany
	fmt.Println(strings.Compare("", ""))
	fmt.Println(strings.Compare("", " ")) // Space is less
}
