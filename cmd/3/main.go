package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

var (
	memoryRe = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	memoryReExt = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\))`)
)

func main() {
	memory := strings.Join(fwk.ReadInputLines(), "\n")

	part1(memory)
	part2(memory)
}

func part1(mem string) {
	var sum uint64 = 0
	for _, m := range memoryRe.FindAllStringSubmatch(mem, -1) {
		arg1, _ := strconv.ParseUint(m[1], 10, 64)
		arg2, _ := strconv.ParseUint(m[2], 10, 64)
		sum += arg1 * arg2
	}

	fwk.Solution(1, sum)
}

func part2(mem string) {
	var sum uint64 = 0
	ignoreMul := false

	for _, m := range memoryReExt.FindAllStringSubmatch(mem, -1) {
		if strings.HasPrefix(m[0], "mul(") {
			if ignoreMul {
				continue
			}
			arg1, _ := strconv.ParseUint(m[2], 10, 64)
			arg2, _ := strconv.ParseUint(m[3], 10, 64)
			sum += arg1 * arg2
		} else if strings.HasPrefix(m[0], "do(") {
			ignoreMul = false
		} else if strings.HasPrefix(m[0], "don't(") {
			ignoreMul = true
		}
	}

	fwk.Solution(2, sum)
}
