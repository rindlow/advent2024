package utils

import (
	"bufio"
	"iter"
	"log"
	"os"
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

func ReadLinesIter(filename string) iter.Seq[string] {
	return func(yield func(string) bool) {
		fileHandle, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Failed to open %s: %q", filename, err)
		}
		defer fileHandle.Close()

		scanner := bufio.NewScanner(fileHandle)
		for scanner.Scan() {
			yield(scanner.Text())
		}
	}
}
