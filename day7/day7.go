package day7

import (
	"math/big"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type Checker func(*big.Int, *big.Int, []*big.Int) bool

func existsAnswerSumMul(answer *big.Int, first *big.Int, rest []*big.Int) bool {
	if len(rest) == 0 {
		return answer.Cmp(first) == 0
	}
	operand := rest[0]
	rest = rest[1:]

	sum := new(big.Int).Add(first, operand)
	if sum.Cmp(answer) <= 0 && existsAnswerSumMul(answer, sum, rest) {
		return true
	}
	prod := new(big.Int).Mul(first, operand)
	return prod.Cmp(answer) <= 0 && existsAnswerSumMul(answer, prod, rest)
}

func existsAnswerSumMulConc(answer *big.Int, first *big.Int, rest []*big.Int) bool {
	if len(rest) == 0 {
		return answer.Cmp(first) == 0
	}
	operand := rest[0]
	rest = rest[1:]

	sum := new(big.Int).Add(first, operand)
	if sum.Cmp(answer) <= 0 && existsAnswerSumMulConc(answer, sum, rest) {
		return true
	}
	prod := new(big.Int).Mul(first, operand)
	if prod.Cmp(answer) <= 0 && existsAnswerSumMulConc(answer, prod, rest) {
		return true
	}
	conc := new(big.Int)
	conc.SetString(first.String()+operand.String(), 10)
	return conc.Cmp(answer) <= 0 && existsAnswerSumMulConc(answer, conc, rest)
}

func sumCalibration(filename string, checker Checker) (sum *big.Int) {
	sum = big.NewInt(0)
	for _, line := range utils.ReadLines(filename) {
		colon := strings.Split(line, ": ")
		answer := new(big.Int)
		answer.SetString(colon[0], 10)
		operands := []*big.Int{}
		for _, digits := range strings.Split(colon[1], " ") {
			num := new(big.Int)
			num.SetString(digits, 10)
			operands = append(operands, num)
		}
		if checker(answer, operands[0], operands[1:]) {
			sum.Add(sum, answer)
		}
	}
	return
}

func Part1(filename string) string {
	return sumCalibration(filename, existsAnswerSumMul).String()
}

func Part2(filename string) string {
	return sumCalibration(filename, existsAnswerSumMulConc).String()
}
