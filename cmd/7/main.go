package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

type Equation struct {
	Expected uint64
	Args     []int64
}

type Operation func(uint64, uint64) uint64

var (
	operationAdd = func(arg1, arg2 uint64) uint64 {
		return arg1 + arg2
	}
	operationMul = func(arg1, arg2 uint64) uint64 {
		return arg1 * arg2
	}
	operationConcat = func(arg1, arg2 uint64) uint64 {
		mul := uint64(math.Pow(10., math.Ceil(math.Log10(math.Max(2, float64(arg2))))))
		return arg1*mul + arg2
	}

	operationsP1 = []Operation{operationAdd, operationMul}

	operationsP2 = []Operation{operationAdd, operationMul, operationConcat}
)

func main() {
	input := ReadInput()

	part1(input)
	part2(input)
}

func part1(equations []Equation) {
	var sum uint64

	for _, equation := range equations {
		if CountPossibleSolutions(equation, operationsP1, uint64(equation.Args[0]), 1) > 0 {
			sum += equation.Expected
		}
	}

	fwk.Solution(1, sum)
}

func part2(equations []Equation) {
	var sum uint64

	for _, equation := range equations {
		if CountPossibleSolutions(equation, operationsP2, uint64(equation.Args[0]), 1) > 0 {
			sum += equation.Expected
		}
	}

	fwk.Solution(2, sum)
}

func CountPossibleSolutions(equation Equation, operations []Operation, sum uint64, pos int) uint64 {
	if pos >= len(equation.Args) {
		if sum == equation.Expected {
			return 1
		}
		return 0
	}

	var count uint64
	arg := uint64(equation.Args[pos])
	for _, operation := range operations {
		lsum := operation(sum, arg)
		count += CountPossibleSolutions(equation, operations, lsum, pos+1)
	}

	return count
}

func ReadInput() []Equation {
	var equations []Equation

	for _, line := range fwk.ReadInputLines() {
		p := strings.Split(line, ": ")
		val, _ := strconv.ParseUint(p[0], 10, 64)
		var args []int64

		for _, arg := range strings.Split(p[1], " ") {
			val, _ := strconv.ParseInt(arg, 10, 64)
			args = append(args, val)
		}

		equations = append(equations, Equation{Expected: val, Args: args})
	}

	return equations
}
