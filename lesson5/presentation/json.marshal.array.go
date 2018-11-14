package main

import (
	"encoding/json"
	"fmt"
)

type Example struct {
	Arr   [2]string
	Slice []int  // null
	Byte  []byte // base64-encoded
}

func main() {
	m := Example{
		Arr:  [2]string{"one", "two"},
		Byte: []byte("some text"),
	}
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
}
