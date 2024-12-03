package fwk

import (
	"bytes"
	"fmt"
)

func StringifyRunesLines(pattern [][]rune) string {
	var buf bytes.Buffer
	for _, line := range pattern {
		buf.WriteString((string(line)))
		buf.WriteString("\n")
	}
	return buf.String()
}

func PrintRunesLines(pattern [][]rune) {
	fmt.Println(StringifyRunesLines(pattern))
}
