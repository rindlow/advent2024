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

func StringToNumbers(s string) (numbers []int) {
	for _, num := range strings.Split(s, " ") {
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
