package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "ping"
		fmt.Println("sent signal")
	}()

	p := <-ch
	fmt.Println("received signal :", p)

	time.Sleep(time.Second)
}
