package main

import (
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

const (
	MIN_CHANGE = 1
	MAX_CHANGE = 3
)

func main() {
	reports := parseReports()

	part1(reports)
	part2(reports)
}

func parseReports() [][]int {
	reportsLines := fwk.ReadInputLines()

	reports := make([][]int, len(reportsLines))
	for i, line := range reportsLines {
		levels := strings.Split(line, " ")

		reports[i] = make([]int, len(levels))
		for j, levelStr := range levels {
			val, _ := strconv.Atoi(levelStr)
			reports[i][j] = val
		}
	}

	return reports
}

func part1(reports [][]int) {
	var sum uint64 = 0

	for _, report := range reports {
		if isValidReport(report) {
			sum++
		}
	}

	fwk.Solution(1, sum)
}

func part2(reports [][]int) {
	var sum uint64 = 0

	for _, report := range reports {
		unsafe := countUnsafeLevels(report)

		if len(unsafe) == 0 {
			sum++
			continue
		}

		for _, unsafePos := range unsafe {
			found := false
			for offset := -1; offset <= 1; offset++ {
				removePos := unsafePos + offset
				if isValidReport(createModifiedReport(report, removePos)) {
					found = true
					break
				}
			}

			if found {
				sum++
				break
			}
		}
	}

	fwk.Solution(2, sum)
}

func createModifiedReport(report []int, removePos int) []int {
	if removePos < 0 || removePos >= len(report) {
		return report
	}

	reportMod := make([]int, len(report)-1)

	j := 0
	for i := 0; i < len(report); i++ {
		if i == removePos {
			continue
		}

		reportMod[j] = report[i]
		j++
	}

	return reportMod
}

func isValidReport(report []int) bool {
	return len(countUnsafeLevels(report)) == 0
}

func countUnsafeLevels(report []int) []int {
	var unsafe []int
	direction := report[0] < report[1]

	for i := 0; i < len(report)-1; i++ {
		diff := fwk.Abs(report[i+1] - report[i])
		if diff < MIN_CHANGE || diff > MAX_CHANGE {
			unsafe = append(unsafe, i)
			continue
		}

		if direction != (report[i] < report[i+1]) {
			unsafe = append(unsafe, i)
			continue
		}
	}
	return unsafe
}
