package main

import (
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

type PuzzleInput struct {
	PageOrderingRules [][2]int
	PageToProduce     [][]int
}

func main() {
	input := ReadInput()

	part1(input)
	part2(input)
}

func part1(input PuzzleInput) {
	var sum int = 0

	for _, update := range input.PageToProduce {
		if input.IsInCorrectOrder(update) {
			// fmt.Printf("Update is correct: %v.\n", update)
			sum += update[len(update)/2]
		}
	}

	fwk.Solution(1, sum)
}

func part2(input PuzzleInput) {
	var sum int = 0

	for _, update := range input.PageToProduce {
		if !input.IsInCorrectOrder(update) {
			sorted := input.Sort(update)

			sum += update[len(sorted)/2]
		}
	}

	fwk.Solution(2, sum)
}

func (input *PuzzleInput) IsInCorrectOrder(update []int) bool {
	for i, _ := range update {
		if input.IsValueInCorrectOrder(update, i) {
		} else {
			return false
		}
	}

	return true
}

func (input *PuzzleInput) IsValueInCorrectOrder(update []int, i int) bool {
	val := update[i]
	before := update[:i]
	after := update[i+1:]

	for range input.PageOrderingRules {
		for _, rule := range input.PageOrderingRules {
			if val == rule[0] {
				afterValue := rule[1]
				// check values before
				for _, v := range before {
					if v == afterValue {
						return false
					}
				}
			} else if val == rule[1] {
				// check values after
				beforeValue := rule[0]
				for _, v := range after {
					if v == beforeValue {
						// fmt.Printf("%d violates rule %v.\n", v, rule)
						return false
					}
				}
			}
		}
	}

	return true
}

func (input *PuzzleInput) Sort(queue []int) []int {
	foundSolution := false
	for !foundSolution {

		foundSolution = true
		for i := 0; i < len(queue); i++ {
			val := queue[i]

			for _, rule := range input.PageOrderingRules {
				if val == rule[0] {
					afterValue := rule[1]
					// check values before
					for j := 0; j < i; j++ {
						v := queue[j]
						if v == afterValue {
							foundSolution = false

							queue[j] = val
							queue[i] = v
							break
						}
					}
				}
				if !foundSolution {
					break
				}
			}
		}
	}
	return queue
}

func ReadInput() PuzzleInput {
	line := strings.TrimSpace(fwk.ReadInput(fwk.GetDataFilename(1)))

	parts := strings.Split(line, "\n\n")

	var rules [][2]int
	for _, line := range strings.Split(parts[0], "\n") {
		ruleParts := strings.Split(line, "|")

		a, _ := strconv.Atoi(ruleParts[0])
		b, _ := strconv.Atoi(ruleParts[1])

		rules = append(rules, [2]int{a, b})
	}

	var produce [][]int
	for _, line := range strings.Split(parts[1], "\n") {
		var update []int
		for _, orderStr := range strings.Split(line, ",") {
			val, _ := strconv.Atoi(orderStr)
			update = append(update, val)
		}
		produce = append(produce, update)
	}

	return PuzzleInput{rules, produce}
}
