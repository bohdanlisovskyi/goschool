package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 9; i++ {
		go func(i int) {
			fmt.Printf("Teacher %d\n", i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}
