package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

type ClawMachine struct {
	Buttons [2][]uint64
	Prize   []uint64
}

const (
	MAX_PUSHES  = 100
	ADDITION_P2 = 10000000000000
)

var PUSH_PRICES = [2]uint64{
	3,
	1,
}

func main() {
	input := ParseInput()

	part1(input)
	part2(input)
}

func part1(input []ClawMachine) {
	var sum uint64 = 0

	for _, machine := range input {
		sum += CountRequiredTokensP1(machine)
	}

	fwk.Solution(1, sum)
}

func part2(input []ClawMachine) {
	var sum uint64 = 0

	p2Input := slices.Clone(input)
	for i, m := range p2Input {
		p2Input[i].Prize = []uint64{m.Prize[0] + ADDITION_P2, m.Prize[1] + ADDITION_P2}
	}

	out := fwk.ComputeParallel(p2Input, CountRequiredTokensP2)
	for tokens := range out {
		sum += tokens
	}

	fwk.Solution(2, sum)
}

func CountRequiredTokensP1(machine ClawMachine) uint64 {
	bestPrice := uint64(0)

	for a := uint64(0); a < MAX_PUSHES; a++ {
		for b := uint64(0); b < MAX_PUSHES; b++ {
			x := a*machine.Buttons[0][0] + b*machine.Buttons[1][0]
			y := a*machine.Buttons[0][1] + b*machine.Buttons[1][1]

			if x == machine.Prize[0] && y == machine.Prize[1] {
				price := a*PUSH_PRICES[0] + b*PUSH_PRICES[1]
				if bestPrice == 0 || bestPrice > price {
					bestPrice = price
				}
			} else if x > machine.Prize[0] || y > machine.Prize[1] {
				break
			}
		}
	}

	return bestPrice
}

func CountRequiredTokensP2(machine ClawMachine) uint64 {
	solution := SolveWithZ3(machine)
	if solution == nil {
		return 0
	}

	return solution[0]*PUSH_PRICES[0] + solution[1]*PUSH_PRICES[1]
}

var valRe = regexp.MustCompile(`\(define-fun (\w+) \(\) Int\s+(\d+)\)`)

func SolveWithZ3(m ClawMachine) []uint64 {
	file := WriteZ3File(m)
	defer os.Remove(file)

	cmd := exec.Command("z3", file)
	out, err := cmd.Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return nil
		}
		panic(err)
	}

	res := make([]uint64, 2)
	for _, m := range valRe.FindAllStringSubmatch(string(out), -1) {
		i := 0
		if m[1] == "BPushes" {
			i = 1
		}
		val, _ := strconv.ParseUint(m[2], 10, 64)
		res[i] = val
	}

	return res
}

func WriteZ3File(m ClawMachine) string {
	file, err := os.CreateTemp("", "aoc-2024-day13.*.smt2")
	check(err)
	defer file.Close()

	fmt.Fprintf(file, `
(declare-const APushes Int)
(assert (>= APushes 0))

(declare-const BPushes Int)
(assert (>= BPushes 0))

(assert 
	(=
		%d
		(+
			(* APushes %d)
			(* BPushes %d)
		)
	)
)

(assert 
	(=
		%d
		(+
			(* APushes %d)
			(* BPushes %d)
		)
	)
)

(minimize
	(+
		(* APushes %d)
		(* BPushes %d)
	)
)

(check-sat)
(get-model)
`,
		m.Prize[0], m.Buttons[0][0], m.Buttons[1][0],
		m.Prize[1], m.Buttons[0][1], m.Buttons[1][1],
		PUSH_PRICES[0], PUSH_PRICES[1],
	)

	return file.Name()
}

var (
	buttonRe = regexp.MustCompile(`(?m)^Button \w+: X\+(\d+), Y\+(\d+)$`)
	prizeRe  = regexp.MustCompile(`(?m)^Prize: X=(\d+), Y=(\d+)$`)
)

func ParseInput() []ClawMachine {
	var machines []ClawMachine
	content := fwk.ReadInput(fwk.GetDataFilename(1))
	for _, section := range strings.Split(content, "\n\n") {
		var machine ClawMachine

		parts := strings.Split(section, "\n")
		for i, machineStr := range parts[:2] {
			match := buttonRe.FindStringSubmatch(machineStr)
			x, _ := strconv.ParseUint(match[1], 10, 64)
			y, _ := strconv.ParseUint(match[2], 10, 64)
			machine.Buttons[i] = []uint64{x, y}
		}

		match := prizeRe.FindStringSubmatch(parts[2])
		x, _ := strconv.ParseUint(match[1], 10, 64)
		y, _ := strconv.ParseUint(match[2], 10, 64)
		machine.Prize = []uint64{x, y}

		machines = append(machines, machine)
	}

	return machines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
