package main

import (
	"fmt"
	"os"
)

func main() {
	home := os.Getenv("USER")
	fmt.Print(home)
	os.NewFile()
	os.Create()
}
