package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang")
	go printName("Teacher")
	time.Sleep(time.Second * 2)
}

func printName(name string) {
	fmt.Println(name)
}
