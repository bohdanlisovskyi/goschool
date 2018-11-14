package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("time now: %s\n", t)

	t1 := t.AddDate(1, 4, 5)
	fmt.Printf("time now: %s\n", t1)
}
