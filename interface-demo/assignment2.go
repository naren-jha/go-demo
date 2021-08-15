package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	// os
	// func Open(name string) (*File, error)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file) // this is some text.
}
