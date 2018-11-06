package main

import "fmt"

func main() {
	go printName("Teacher")
}

func printName(name string) {
	fmt.Println(name)
}
