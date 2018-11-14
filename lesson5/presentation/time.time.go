package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Year: %d\n", t.Year())
	fmt.Printf("Month: %d\n", t.Month())
	fmt.Printf("Day: %d\n", t.Day())
	fmt.Printf("Hour: %d\n", t.Hour())
	fmt.Printf("Minute: %d\n", t.Minute())
	fmt.Printf("Second: %d\n", t.Second())
	fmt.Printf("Nanosecond: %d\n", t.Nanosecond())
}
