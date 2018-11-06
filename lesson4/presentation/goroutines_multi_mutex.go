package main

import (
	"sync"
)

var mutex sync.Mutex
var counter = 0

func main() {

	go func() {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}()
}
