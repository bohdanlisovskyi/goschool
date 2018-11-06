package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Teacher %d\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
