package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileName := string(os.Args[1])

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
