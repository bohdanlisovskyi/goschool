package main

import (
	"sync"
)

var counter = 0

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func(i int) {
			counter++
			wg.Done()
		}(i)
	}

	wg.Wait()
}
