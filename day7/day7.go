package day7

import (
	"math/big"
	"strings"

	"github.com/rindlow/aoc-utils"
)

type Combiner func([]*big.Int) []*big.Int

func addOrMul(operands []*big.Int) []*big.Int {
	if len(operands) == 1 {
		return operands
	}
	sum := new(big.Int)
	sum.Add(operands[0], operands[1])
	prod := new(big.Int)
	prod.Mul(operands[0], operands[1])
	if len(operands) == 2 {
		return []*big.Int{sum, prod}
	}
	summed := addOrMul(append([]*big.Int{sum}, operands[2:]...))
	prodded := addOrMul(append([]*big.Int{prod}, operands[2:]...))
	return append(summed, prodded...)
}

func addMulOrConcat(operands []*big.Int) []*big.Int {
	if len(operands) == 1 {
		return operands
	}
	sum := new(big.Int)
	sum.Add(operands[0], operands[1])
	prod := new(big.Int)
	prod.Mul(operands[0], operands[1])
	conc := new(big.Int)
	conc.SetString(operands[0].String()+operands[1].String(), 10)
	if len(operands) == 2 {
		return []*big.Int{sum, prod, conc}
	}
	summed := addMulOrConcat(append([]*big.Int{sum}, operands[2:]...))
	prodded := addMulOrConcat(append([]*big.Int{prod}, operands[2:]...))
	conced := addMulOrConcat(append([]*big.Int{conc}, operands[2:]...))
	return append(append(summed, prodded...), conced...)
}

func testCalibration(answer *big.Int, operands []*big.Int, combiner Combiner) bool {
	for _, calc := range combiner(operands) {
		if answer.Cmp(calc) == 0 {
			return true
		}
	}
	return false
}

func sumCalibration(filename string, combiner Combiner) (sum *big.Int) {
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
		if testCalibration(answer, operands, combiner) {
			sum.Add(sum, answer)
		}
	}
	return
}

func Part1(filename string) string {
	return sumCalibration(filename, addOrMul).String()
}

func Part2(filename string) string {
	return sumCalibration(filename, addMulOrConcat).String()
}
