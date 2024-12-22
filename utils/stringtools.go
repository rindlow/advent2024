package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) (lines []string) {
	fileHandle, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s: %q", filename, err)
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func ReadLinesAsInt(filename string) (lines []int) {
	fileHandle, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s: %q", filename, err)
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Atoi %s: %q", line, err)
		}
		lines = append(lines, num)
	}
	return
}

func StringToNumbersWithDelimiter(s string, delimiter string) (numbers []int) {
	for _, num := range strings.Split(s, delimiter) {
		if num == "" {
			continue
		}
		n, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			log.Fatalf("'%s': %s", num, err)
		}
		numbers = append(numbers, n)
	}
	return
}

func StringToNumbers(s string) (numbers []int) {
	return StringToNumbersWithDelimiter(s, " ")
}
