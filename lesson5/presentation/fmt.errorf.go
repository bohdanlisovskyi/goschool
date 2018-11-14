package main

import "fmt"

func main() {
	err := fmt.Errorf("some error messagd with status code %d\n", 401)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%T\n", err)
}
