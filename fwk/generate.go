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
