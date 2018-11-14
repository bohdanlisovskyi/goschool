package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("/Users/bohdan/projects/gomaytech/src/golangschool/lesson5/presentation/")

	r, err := f.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", r)
}
