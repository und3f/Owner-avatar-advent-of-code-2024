package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

const (
	SIDE    = 1 + 70
	StepsP1 = 1024
)

func main() {
	grid, restBytes := ParseInput()
	Part1(grid)
	Part2(grid, restBytes)
}

func Part1(grid [][]rune) {
	graph := fwk.NewGridGraph(grid)
	bfs := fwk.NewBFS(graph, 0)

	fwk.Solution(1, len(bfs.PathTo(graph.Len()-1))-1)
}

func Part2(grid [][]rune, bytesToFail [][]int) {
	for _, pos := range bytesToFail {
		grid[pos[0]][pos[1]] = '#'
		graph := fwk.NewGridGraph(grid)
		v := fwk.HashVect(grid, pos)

		bfs := fwk.NewBFS(graph, 0)
		path := bfs.PathTo(graph.Len() - 1)
		if path == nil {
			y := v / SIDE
			x := v % SIDE
			fwk.Solution(2, fmt.Sprintf("%d,%d", x, y))
			return
		}
	}

}

func ParseInput() ([][]rune, [][]int) {
	lines := fwk.ReadInputLines()
	input := make([][]int, len(lines))

	for i, l := range lines {
		vec := make([]int, 2)

		for i, str := range strings.Split(l, ",") {
			num, _ := strconv.Atoi(str)
			vec[i] = num
		}
		slices.Reverse(vec)
		input[i] = vec
	}

	grid := make([][]rune, SIDE)
	for i := range grid {
		grid[i] = make([]rune, SIDE)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for i, pos := range input {
		if i >= StepsP1 {
			break
		}
		grid[pos[0]][pos[1]] = '#'
	}

	return grid, input[StepsP1:]
}
