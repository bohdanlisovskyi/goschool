package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat("FileOrDir")
	if err != nil {
		fmt.Println(err)
		return
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		fmt.Println("directory")
	case mode.IsRegular():
		// do file stuff
		fmt.Println("file")
	}
}
