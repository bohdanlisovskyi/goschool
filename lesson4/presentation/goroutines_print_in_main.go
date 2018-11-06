package main

import "fmt"

func main() {
	fmt.Println("Golang")
	go printName("Teacher")
}

func printName(name string) {
	fmt.Println(name)
}
