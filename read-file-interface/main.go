package main

import (
	"io"
	"os"
)

func main() {
	myFile := os.Args[1]

	fileBody, err := os.Open(myFile)
	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, fileBody)
}
