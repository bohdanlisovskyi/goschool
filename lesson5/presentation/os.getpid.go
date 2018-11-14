package main

import (
	"fmt"
	"os"
)

func main() {
	processID := os.Getpid()
	fmt.Print(processID)
}
