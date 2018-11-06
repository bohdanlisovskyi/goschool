package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	wg.Add(2)
	for i := 0; i < 9; i++ {
		go func(i int) {
			fmt.Printf("Teacher %d\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
