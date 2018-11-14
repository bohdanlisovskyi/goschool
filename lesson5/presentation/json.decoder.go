package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	jsonStream := `{"Name": "Ed", "Text": "Knock knock."}`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	var m Message
	dec.Decode(&m)
	fmt.Printf("%s: %s\n", m.Name, m.Text)
}
