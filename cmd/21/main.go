package main

import "github.com/und3f/aoc/2024/fwk"

var keypad = [][]rune{
	[]rune{'7', '8', '9'},
	[]rune{'4', '5', '6'},
	[]rune{'1', '2', '3'},
	[]rune{'#', '0', 'A'},
}

var directionalKeypad = [][]rune{
	[]rune{'#', '^', 'A'},
	[]rune{'<', 'v', '>'},
}

func main() {
	codes := ParseInput()
	Part1(codes)
}

func Part1(codes [][]rune) {
	fwk.GridGraph(keypad)
	fwk.Solution(1, codes)
}

type Keypad struct {
	graph Graph
	pos   int
}

func NewKeypad(keypad [][]rune) *Keypad {
	return &Keypad{keypad, fwk.HashVect(keypad, 'A')}
}

func ParseInput() [][]rune {
	lines := fwk.ReadInputLines()
	codes := make([][]rune, len(lines))
	for i, line := range lines {
		codes[i] = make([]rune, len(line))
		for j, code := range line {
			codes[i][j] = code
		}
	}

	return codes
}

func FindSymbol(maze [][]rune, sym rune) []int {
	for i, row := range maze {
		for j, v := range row {
			if v == sym {
				return []int{i, j}
			}
		}
	}

	return nil
}
