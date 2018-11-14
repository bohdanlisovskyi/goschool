package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("time now: %s\n", t)
	fmt.Printf("time now: %s\n", t.Format(time.UnixDate))
	fmt.Printf("time now: %s\n", t.Format("2 Jan 2006"))
}
