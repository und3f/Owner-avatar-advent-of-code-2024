package main

import (
	"github.com/und3f/aoc/2024/fwk"
)

func main() {
	input := fwk.ReadInputRunesLines()

	part1(input)
	part2(input)
}

func part1(input [][]rune) {
	antinodes := calcAntinodes(input)
	sum := calcAntinodesCount(antinodes)
	fwk.Solution(1, sum)
}

func part2(input [][]rune) {
	antinodes := calcAntinodesP2(input)
	sum := calcAntinodesCount(antinodes)
	fwk.Solution(2, sum)
}

func calcAntinodesCount(antinodes [][]rune) int {
	var sum int
	for _, row := range antinodes {
		for _, v := range row {
			if v == '#' {
				sum++
			}
		}
	}
	return sum
}

func calcAntinodes(amap [][]rune) [][]rune {
	antennas := buildAntennasMap(amap)

	antinodes := buildAntinodes(amap)

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			p1 := positions[i]
			for j := i + 1; j < len(positions); j++ {
				p2 := positions[j]
				vec := fwk.SubVec(p1, p2)
				a := [][]int{fwk.AddVec(p1, vec), fwk.SubVec(p2, vec)}

				for _, antifrequency := range a {
					if fwk.IsVecIn2DBounds(antinodes, antifrequency) {
						antinodes[antifrequency[0]][antifrequency[1]] = '#'
					}
				}
			}
		}
	}

	return antinodes
}

func calcAntinodesP2(amap [][]rune) [][]rune {
	antennas := buildAntennasMap(amap)

	antinodes := buildAntinodes(amap)

	for _, positions := range antennas {
		for i := 0; i < len(positions)-1; i++ {
			p1 := positions[i]
			for j := i + 1; j < len(positions); j++ {
				p2 := positions[j]
				vec := fwk.SubVec(p1, p2)

				for antifrequency := fwk.AddVec(p1, vec); fwk.IsVecIn2DBounds(antinodes, antifrequency); antifrequency = fwk.AddVec(antifrequency, vec) {
					antinodes[antifrequency[0]][antifrequency[1]] = '#'
				}

				for antifrequency := fwk.SubVec(p2, vec); fwk.IsVecIn2DBounds(antinodes, antifrequency); antifrequency = fwk.SubVec(antifrequency, vec) {
					antinodes[antifrequency[0]][antifrequency[1]] = '#'
				}

				antinodes[p1[0]][p1[1]] = '#'
				antinodes[p2[0]][p2[1]] = '#'
			}
		}
	}

	return antinodes
}

func buildAntennasMap(amap [][]rune) map[rune][][]int {
	antennas := make(map[rune][][]int)
	for i, row := range amap {
		for j, v := range row {
			if v != '.' {
				pos := []int{i, j}
				if _, exist := antennas[v]; !exist {
					antennas[v] = [][]int{pos}
				} else {
					antennas[v] = append(antennas[v], pos)
				}
			}
		}
	}
	return antennas
}

func buildAntinodes(amap [][]rune) [][]rune {
	antinodes := make([][]rune, len(amap))
	for i, _ := range amap {
		antinodes[i] = make([]rune, len(amap[i]))
		for j := range antinodes[i] {
			antinodes[i][j] = '.'
		}
	}
	return antinodes
}
