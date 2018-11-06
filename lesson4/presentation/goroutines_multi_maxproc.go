package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)

	for i := 0; i < 9; i++ {
		go func(i int) {
			fmt.Printf("Teacher %d\n", i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
