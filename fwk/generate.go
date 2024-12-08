package fwk

import "bytes"

func GenerateInts(start int, end int) []int {
	if start >= end {
		panic("Invalid ints rage")
	}

	ints := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		ints[i] = i
	}

	return ints
}

func GenerateStringRange(start rune, end rune) string {
	var buf bytes.Buffer

	for i := start; i <= end; i++ {
		buf.WriteRune(i)
	}

	return buf.String()
}

func GenerateRunesLines(height, width int, fill rune) [][]rune {
	board := make([][]rune, height)
	for i, _ := range board {
		board[i] = make([]rune, width)
		for j := range board[i] {
			board[i][j] = fill
		}
	}
	return board
}
