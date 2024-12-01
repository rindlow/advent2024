package utils

import (
	"log"
	"strconv"
	"strings"
)

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
