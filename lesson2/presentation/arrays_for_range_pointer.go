package main

import "fmt"

func main() {

	farm := [3]string{"Cow", "Pig", "Chicken"}
	fmt.Printf("Animal[%s] : ", farm[1])

	for i := range farm {
		farm[1] = "Horse"

		if i == 1 {
			fmt.Printf("v[%s]\n", farm[i])
		}
	}

	fmt.Println(farm)
}
