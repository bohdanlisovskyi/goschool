package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("time now: %s\n", t)

	t1 := t.Add(5 * time.Minute)
	fmt.Printf("time now: %s\n", t1)
}
