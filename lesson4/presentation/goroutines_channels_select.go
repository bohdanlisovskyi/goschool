package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	//START OMIT
	go func() {
		ch1 <- "ping"
	}()
	go func() {
		ch2 <- "pong"
	}()
	go func() {
		for {
			select {
			case <-ch1:
				fmt.Printf("result1 \n")
			case <-ch2:
				fmt.Printf("result2 \n")

			}
		}
	}()
	time.Sleep(time.Second * 2)
	//END OMIT
}
