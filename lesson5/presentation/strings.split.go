package main

import (
	"fmt"
	"strings"
)

func main() {

	maytech := "M-a-y-t-e-c-h"
	maytechSplit := strings.Split(maytech, "-")

	for _, v := range maytechSplit {
		fmt.Println(v)
	}
}
