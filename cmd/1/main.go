package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

var ()

func main() {
	left, right := parseLocationLists()

	part1(left, right)
	part2(left, right)
}

func parseLocationLists() ([]int, []int) {
	lines := fwk.ReadInputLines()

	lists := make([][2]int, len(lines))
	for i, line := range lines {
		values := strings.Split(line, "   ")

		for j, valueStr := range values {
			val, _ := strconv.Atoi(valueStr)
			lists[i][j] = val
		}
	}

	left, right := split(lists)
	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func part1(left []int, right []int) {
	var sum uint64 = 0

	for i, valA := range left {
		valB := right[i]

		sum += uint64(fwk.Abs(valB - valA))
	}

	fwk.Solution(1, sum)
}

func part2(left []int, right []int) {
	var sum uint64 = 0

	for _, value := range left {
		sum += uint64(countOccurancesInSortedSlice(right, value)) * uint64(value)
	}

	fwk.Solution(2, sum)
}

func countOccurancesInSortedSlice(arr []int, value int) int {
	ind := fwk.FindValueInSortedSlice(arr, value)
	if ind == -1 {
		return 0
	}

	return findRightmost(arr, ind) - findLeftmost(arr, ind) + 1
}

func findRightmost(arr []int, ind int) int {
	var i int
	for i = ind; i < len(arr) && arr[i] == arr[ind]; i++ {
	}
	return i - 1
}

func findLeftmost(arr []int, ind int) int {
	var i int
	for i = ind; i >= 0 && arr[i] == arr[ind]; i-- {
	}
	return i + 1
}

func split(lists [][2]int) ([]int, []int) {
	a := make([]int, len(lists))
	b := slices.Clone(a)

	for i, v := range lists {
		a[i] = v[0]
		b[i] = v[1]
	}

	return a, b
}
