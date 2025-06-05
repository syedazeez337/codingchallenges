package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type WcResult struct {
	Lines    int
	Words    int
	Bytes    int
	Chars    int
	Filename string
}

func getWCResult(filename string) (WcResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return WcResult{}, err
	}
	defer file.Close()

	// Read entire file once
	data, err := io.ReadAll(file)
	if err != nil {
		return WcResult{}, err
	}

	text := string(data)
	lines := 0
	words := 0

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		lines++
	}
	wordScanner := bufio.NewScanner(strings.NewReader(text))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		words++
	}

	return WcResult{
		Lines:    lines,
		Words:    words,
		Bytes:    len(data),
		Chars:    len([]rune(text)),
		Filename: filename,
	}, nil
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: ccwc [-c|-l|-w|-m] <filename")
		return
	}

	var option, filename string
	if len(os.Args) == 2 {
		option = "default"
		filename = os.Args[1]
	} else {
		option = os.Args[1]
		filename = os.Args[2]
	}

	result, err := getWCResult(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch option {
	case "-c":
		fmt.Printf("%8d %s\n", result.Bytes, result.Filename)
	case "-l":
		fmt.Printf("%8d %s\n", result.Lines, result.Filename)
	case "-w":
		fmt.Printf("%8d %s\n", result.Words, result.Filename)
	case "-m":
		fmt.Printf("%8d %s\n", result.Words, result.Filename)
	case "default":
		fmt.Printf("%8d %8d %8d %s\n", result.Lines, result.Words, result.Bytes, result.Filename)
	default:
		fmt.Println("Unknown option:", option)
	}
}

/*
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
		fmt.Printf("%8T %s\n", count, filename)

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

	case "-m":
		count, err := countCharacters(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%8d %s\n", count, filename)

	default:
		// fmt.Println("Unknown Option:", option)
		fmt.Printf("%8d %8d %8d %s\n", )
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

func countCharacters(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	runes := []rune(string(data))
	return len(runes), nil
}

*/
