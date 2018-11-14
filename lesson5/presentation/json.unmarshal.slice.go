package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`["one","two"]`)

	var u []string
	err := json.Unmarshal(jsonBlob, &u)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%T %+v", u, u)
}
