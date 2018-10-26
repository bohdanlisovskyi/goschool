package main

import "fmt"

func main() {

	farm := [3]string{"Cow", "Pig", "Chicken"}

	for i := 0; i < len(farm); i++ {
		fmt.Println(farm[i])
	}
}
