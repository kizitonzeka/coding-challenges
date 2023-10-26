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
	bytes := flag.String("c", "", "count bytes in file.")
	words := flag.String("w", "", "count number of words in file.")
	lines := flag.String("l", "", "count number of lines in file.")
	chars := flag.String("m", "", "count number of characters in file.")

	flag.Parse()

	switch true {

	case isFlagPassed("c"):
		totalBytes, err := countBytes(bytes)
		if err != nil {
			log.Panic("failed to count bytes: ", err)
		}
		fmt.Printf("%d\t%s", totalBytes, *bytes)
	case isFlagPassed("w"):
		totalWords, err := countWords(words)
		if err != nil {
			log.Panic("failed to count words: ", err)
		}
		fmt.Printf("%d\t%s", totalWords, *words)
	case isFlagPassed("l"):
		totalLines, err := countLines(lines)
		if err != nil {
			log.Panic("failed to count lines: ", err)
		}
		fmt.Printf("%d\t %s", totalLines, *lines)
	case isFlagPassed("m"):
		totalChars, err := countChars(chars)
		if err != nil {
			log.Panic("failed to count chars: ", err)
		}
		fmt.Printf("%d\t%s", totalChars, *chars)
	default:

		fileName := flag.Arg(0)
		file := &fileName
		totalLines, _ := countLines(file)
		totalWords, _ := countWords(file)
		totalBytes, _ := countBytes(file)
		fmt.Printf("%d\t%d\t%d\t%s", totalLines, totalWords, totalBytes, fileName)
	}

}

func isFlagPassed(n string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == n {
			found = true
		}
	})
	return found
}

func countBytes(flag *string) (int64, error) {
	file, err := os.Open(*flag)
	if err != nil {
		return 0, err
	}
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

func countWords(flag *string) (int64, error) {
	file, err := os.Open(*flag)
	if err != nil {
		return 0, err
	}
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

func countLines(flag *string) (int64, error) {
	file, err := os.Open(*flag)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var count int64
	for scanner.Scan() {
		count++
	}

	return count, nil
}

func countChars(flag *string) (int64, error) {
	file, err := os.Open(*flag)
	if err != nil {
		return 0, err
	}
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
