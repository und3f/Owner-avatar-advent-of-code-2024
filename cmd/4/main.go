package main

import (
	"github.com/und3f/aoc/2024/fwk"
)

var (
	XMAS = []rune("XMAS")
)

func main() {
	wordSearch := fwk.ReadInputRunesLines()

	part1(wordSearch)
	part2(wordSearch)
}

func part1(wordSearch [][]rune) {
	var count uint64 = 0

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch); j++ {
			count += uint64(countWords(wordSearch, i, j, XMAS))
		}
	}

	fwk.Solution(1, count)
}

func part2(wordSearch [][]rune) {
	var count uint64 = 0

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch); j++ {
			if containsXMas(wordSearch, i, j) {
				count++
			}
		}
	}

	fwk.Solution(2, count)
}

func countWords(wordSearch [][]rune, i, j int, expected []rune) int {
	var count int

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			if containsWordInDirection(wordSearch, i, j, di, dj, expected) {
				count++
			}

		}
	}

	return count
}

func containsWordInDirection(wordSearch [][]rune, i, j, di, dj int, expected []rune) bool {
	if len(expected) == 0 {
		return true
	}

	if wordSearch[i][j] != expected[0] {
		return false
	}
	if len(expected) == 1 {
		return true
	}

	ni := i + di
	nj := j + dj

	if ni < 0 || ni >= len(wordSearch) {
		return false
	}
	if nj < 0 || nj >= len(wordSearch[ni]) {
		return false
	}

	if containsWordInDirection(wordSearch, ni, nj, di, dj, expected[1:]) {
		return true
	}

	return false
}

func containsXMas(wordSearch [][]rune, i, j int) bool {
	if !(i >= 1 && i < len(wordSearch)-1 &&
		j >= 1 && j < len(wordSearch[i])-1) {
		return false
	}

	return wordSearch[i][j] == 'A' &&
		((wordSearch[i-1][j-1] == 'M' && wordSearch[i+1][j-1] == 'M' &&
			wordSearch[i-1][j+1] == 'S' && wordSearch[i+1][j+1] == 'S') ||
			(wordSearch[i-1][j-1] == 'M' && wordSearch[i-1][j+1] == 'M' &&
				wordSearch[i+1][j-1] == 'S' && wordSearch[i+1][j+1] == 'S') ||
			(wordSearch[i-1][j-1] == 'S' && wordSearch[i+1][j-1] == 'S' &&
				wordSearch[i-1][j+1] == 'M' && wordSearch[i+1][j+1] == 'M') ||
			(wordSearch[i-1][j-1] == 'S' && wordSearch[i-1][j+1] == 'S' &&
				wordSearch[i+1][j-1] == 'M' && wordSearch[i+1][j+1] == 'M'))
}
