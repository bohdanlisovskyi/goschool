package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = f.Chmod(0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
