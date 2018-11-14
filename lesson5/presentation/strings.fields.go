package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "Secure, compliant, simple cloud file sharing and data storage"
	testArray := strings.Fields(testString)
	for _, v := range testArray {
		fmt.Println(v)
	}
}
