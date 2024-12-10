package main

import (
	queue "github.com/Jcowwell/go-algorithm-club/Queue"
	"github.com/und3f/aoc/2024/fwk"
	"github.com/und3f/aoc/2024/fwk/twoD"
)

func main() {
	input := fwk.ReadInputRunesLines()

	part1(input)
	part2(input)
}

func part1(input [][]rune) {
	sum := 0

	for _, startPos := range findStartPos(input) {
		sum += fwk.CountRunes(evaluateTrainhead(input, startPos), '9')
	}

	fwk.Solution(1, sum)
}

func part2(input [][]rune) {
	sum := 0

	for _, startPos := range findStartPos(input) {
		sum += evaluateTrainheadCountTrails(input, startPos)
	}

	fwk.Solution(1, sum)
}

func evaluateTrainhead(input [][]rune, startPos [2]int) [][]rune {
	reachable := fwk.GenerateRunesLines(len(input), len(input[0]), '.')
	q := queue.Queue[[2]int]{}
	q.Enqueue(startPos)
	for !q.IsEmpty() {
		p, _ := q.Dequeue()
		pV := input[p[0]][p[1]]
		// fmt.Println(p, pV)

		for _, direction := range [][]int{
			twoD.DirectionNorth,
			twoD.DirectionSouth,
			twoD.DirectionEast,
			twoD.DirectionWest,
		} {
			wSlice := fwk.AddVec(p[:], direction)
			if !fwk.IsVecIn2DBounds(input, wSlice) {
				continue
			}

			w := [2]int{wSlice[0], wSlice[1]}
			wV := input[w[0]][w[1]]

			if wV != pV+1 {
				continue
			}

			reachable[w[0]][w[1]] = wV
			q.Enqueue(w)
		}
	}

	return reachable
}

func evaluateTrainheadCountTrails(input [][]rune, startPos [2]int) int {
	trails := make([]int, 10)

	q := queue.Queue[[2]int]{}
	q.Enqueue(startPos)
	for !q.IsEmpty() {
		p, _ := q.Dequeue()
		pV := input[p[0]][p[1]]
		trails[pV-'0']++

		for _, direction := range [][]int{
			twoD.DirectionNorth,
			twoD.DirectionSouth,
			twoD.DirectionEast,
			twoD.DirectionWest,
		} {
			wSlice := fwk.AddVec(p[:], direction)
			if !fwk.IsVecIn2DBounds(input, wSlice) {
				continue
			}

			w := [2]int{wSlice[0], wSlice[1]}
			wV := input[w[0]][w[1]]

			if wV != pV+1 {
				continue
			}

			q.Enqueue(w)
		}
	}

	return trails[len(trails)-1]
}

func findStartPos(input [][]rune) [][2]int {
	var starts [][2]int
	for i, row := range input {
		for j, v := range row {
			if v == '0' {
				starts = append(starts, [2]int{i, j})
			}
		}
	}
	if len(starts) == 0 {
		panic("Start position not found!")
	}
	return starts
}
