package main

import (
	"fmt"
	"gjp/internal/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jsonparser <file>")
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println("Input JSON:", string(data))

	err = parser.Parse(string(data))
	if err != nil {
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
	os.Exit(0)
}
