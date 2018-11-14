package main

import (
	"fmt"
	"strings"
)

func main() {
	strArr := []string{"M", "a", "y", "t", "e", "c", "h"}
	fmt.Println(strings.Join(strArr, ""))
	fmt.Println(strings.Join(strArr, "-"))
}
