package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Printf("PING: %d", <-ch)

		wg.Done()
	}()

	close(ch)
	wg.Wait()
}
