package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// flags
	_ = flag.String("c", "test.txt", "input file")
	_ = flag.String("w", "test.txt", "input file")
	_ = flag.String("l", "test.txt", "input file")
	charFile := flag.String("m", "test.txt", "input file")

	flag.Parse()

	file, err := os.Open(*charFile)

	if err != nil {
		log.Panic("failed to open file: ", err)
	}
	totalChars, err := countChars(file)
	if err != nil {
		log.Panic("failed to count bytes: ", err)
	}
	fmt.Println(totalChars)

}

func countBytes(file *os.File) (int64, error) {
	defer file.Close()
	var count int64
	buffer := make([]byte, 1024)
	for {
		b, err := file.Read(buffer)
		if err == io.EOF {
			return count, nil
		} else if err != nil {
			return 0, err
		}
		count += int64(b)
	}
}

func countWords(file *os.File) (int64, error) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	var count int64
	for scanner.Scan() {
		count++
	}

	return count, nil
}

func countLines(file *os.File) (int64, error) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count int64
	for scanner.Scan() {
		count++
	}

	return count, nil
}

func countChars(file *os.File) (int64, error) {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	var count int64
	for scanner.Scan() {
		count++
	}

	return count, nil
}
