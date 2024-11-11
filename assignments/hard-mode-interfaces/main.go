package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, error := os.Open(fileName)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	defer file.Close()

	content, readError := io.ReadAll(file)
	if readError != nil {
		fmt.Println("Error:", readError)
		os.Exit(1)
	}

	fmt.Println(string(content))
}
