package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("received signal :", p)
	}()

	time.Sleep(500 * time.Millisecond)
	ch <- "ping"
	fmt.Println("sent signal")
}
