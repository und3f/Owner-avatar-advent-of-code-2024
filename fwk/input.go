package fwk

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func basename(file string) string {
	return strings.TrimSuffix(path.Base(file), path.Ext(file))
}

func lastDirName(file string) string {
	return path.Base(strings.TrimSuffix(file, path.Base(file)))
}

const inputPrefix = "./resources"

func GetDataFilename(depth int) string {
	_, file, _, ok := runtime.Caller(depth)
	if !ok {
		panic("Failed to obtain caller file")
	}

	return path.Join(inputPrefix, fmt.Sprintf("%s%s%s", "day", lastDirName(file), ".txt"))
}

func ReadInput(inputFilename string) string {
	if len(inputFilename) == 0 {
		inputFilename = GetDataFilename(2)
	}

	data, err := os.ReadFile(inputFilename)
	if err != nil {
		log.Fatalf("Failed to read input data from %s: %v", inputFilename, err)
	}

	return string(data)
}

func ReadInputLines() []string {
	content := ReadInput(GetDataFilename(2))
	return strings.Split(strings.TrimSpace(content), "\n")
}

func ReadInputRunesLines() [][]rune {
	lines := strings.Split(ReadInput(GetDataFilename(2)), "\n")

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	runes := make([][]rune, len(lines))
	for i := range lines {
		runes[i] = []rune(lines[i])
	}

	return runes
}
