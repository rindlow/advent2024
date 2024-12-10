package day7

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type Checker func(int, []int) bool

func existsAnswerSumMul(answer int, operands []int) bool {
	if len(operands) == 1 {
		return answer == operands[0]
	}
	first := operands[0]
	if answer%first == 0 && existsAnswerSumMul(answer/first, operands[1:]) {
		return true
	}
	return answer > first && existsAnswerSumMul(answer-first, operands[1:])
}

func existsAnswerSumMulConc(answer int, operands []int) bool {
	if len(operands) == 1 {
		return answer == operands[0]
	}
	first := operands[0]
	if answer%first == 0 && existsAnswerSumMulConc(answer/first, operands[1:]) {
		return true
	}
	strAnswer := strconv.Itoa(answer)
	strFirst := strconv.Itoa(first)
	if answer != first && strings.HasSuffix(strAnswer, strFirst) {
		prefixLen := len(strAnswer) - len(strFirst)
		nextAnswer, err := strconv.Atoi(strAnswer[:prefixLen])
		if err != nil {
			log.Fatalf("Atoi %s: %q", strAnswer[:prefixLen], err)
		}
		if existsAnswerSumMulConc(nextAnswer, operands[1:]) {
			return true
		}
	}
	return answer > first && existsAnswerSumMulConc(answer-first, operands[1:])
}

func sumCalibration(filename string, checker Checker) (sum int) {
	for _, line := range utils.ReadLines(filename) {
		colon := strings.Split(line, ": ")
		answer, err := strconv.Atoi(colon[0])
		if err != nil {
			log.Fatalf("Atoi(%s): %q", colon[0], err)
		}
		operands := []int{}
		for _, digits := range strings.Split(colon[1], " ") {
			num, err := strconv.Atoi(digits)
			if err != nil {
				log.Fatalf("Atoi(%s): %q", digits, err)
			}
			operands = append(operands, num)
		}
		slices.Reverse(operands)
		if checker(answer, operands) {
			sum += answer
		}
	}
	return
}

func Part1(filename string) string {
	return strconv.Itoa(sumCalibration(filename, existsAnswerSumMul))
}

func Part2(filename string) string {
	return strconv.Itoa(sumCalibration(filename, existsAnswerSumMulConc))
}
