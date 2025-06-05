package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ccwc <-c|-l> <filename>")
		return
	}

	option := os.Args[1]
	filename := os.Args[2]

	switch option {
	case "-c":
		count, err := countBytes(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%8d %s\n", count, filename)

	case "-l":
		count, err := countLines(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%8d %s\n", count, filename)

	case "-w":
		count, err := countWords(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%8d %s\n", count, filename)

	default:
		fmt.Println("Unknown Option:", option)
	}

	/*
	if len(os.Args) != 3 || os.Args[1] != "-c" {
		fmt.Println("Usage: ccwc -c <filename>")
		return
	}

	filename := os.Args[2]
	count, err := coutBytes(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%8d %s\n", count, filename)
	*/
}

func countBytes(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return int(stats.Size()), nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}


func countWords(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}